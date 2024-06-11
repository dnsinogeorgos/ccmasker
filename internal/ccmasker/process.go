package ccmasker

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strconv"

	"github.com/theplant/luhn"
)

//easyjson:json
type Message struct {
	Msg []byte `json:"msg"`
}

// ProcessMessage filters the message through regexp filters and returns appropriate response for rsyslog
// The iterations appear wasteful, but there are edge cases which make iterating for
// all possible PAN lengths necessary.
func ProcessMessage(message []byte, filters []filterGroup, numFilter *regexp.Regexp) ([]byte, error) {
	validated := false

	for _, group := range filters {
		// If variable length pattern matches move on
		if group.variable.Match(message) {
			for _, fixedPattern := range group.fixed {
				// If fixed length pattern matches move on
				if fixedPattern.Match(message) {
					matchStrings := fixedPattern.FindAll(message, -1)
					for _, match := range matchStrings {
						// Prepare string for Luhn check
						cleanMatch := numFilter.ReplaceAll(match, []byte{})
						cleanInt, err := strconv.Atoi(string(cleanMatch))
						if err != nil {
							return []byte{}, err
						}
						// Check with Luhn
						if luhn.Valid(cleanInt) {
							validated = true
							message = fixedPattern.ReplaceAllLiteral(message, group.mask)
						}
					}
				}
			}
		}
	}

	// If PAN data isn't found, return empty JSON
	if validated == false {
		return []byte{'{', '}', '\n'}, nil
	}

	// If PAN data is found, wrap to JSON and return
	message = bytes.TrimSuffix(message, []byte{'\n'})
	response, err := json.Marshal(Message{Msg: message})
	if err != nil {
		return []byte{}, err
	}
	return append(response, '\n'), nil
}

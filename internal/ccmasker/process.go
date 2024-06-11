package ccmasker

import (
	"bytes"
	"regexp"
	"strconv"

	"github.com/mailru/easyjson/jwriter"
	"github.com/theplant/luhn"
)

//easyjson:json
type Message struct {
	Msg string `json:"msg,nocopy"`
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
	writer := jwriter.Writer{}
	jsonMessage := Message{Msg: string(message)}
	jsonMessage.MarshalEasyJSON(&writer)
	jsonBytes, err := writer.BuildBytes()
	if err != nil {
		return []byte{}, err
	}
	return append(jsonBytes, '\n'), nil
}

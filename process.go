package main

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	"github.com/theplant/luhn"
)

type Message struct {
	Msg string
}

// ProcessMessage filters the message through regexp filters and returns appropriate response for rsyslog
// The iterations appear wasteful, but there are edge cases which make iterating for
// all possible PAN lengths.
func ProcessMessage(message string, filters []FilterGroup, numFilter *regexp.Regexp) (string, error) {
	validated := false

	for _, group := range filters {
		// If variable length pattern matches move on
		if group.Variable.MatchString(message) {
			for _, fixedPattern := range group.Fixed {
				// If fixed length pattern matches move on
				if fixedPattern.MatchString(message) {
					matchStrings := fixedPattern.FindAllString(message, -1)
					for _, match := range matchStrings {
						// Prepare string for Luhn check
						cleanMatch := numFilter.ReplaceAllString(match, "")
						cleanInt, err := strconv.Atoi(cleanMatch)
						if err != nil {
							return "", err
						}

						// Check with Luhn
						if luhn.Valid(cleanInt) {
							validated = true
							message = fixedPattern.ReplaceAllLiteralString(message, group.Mask)
						}
					}
				}
			}
		}
	}

	// If PAN data isn't found, return empty JSON
	if validated == false {
		return "{}\n", nil
	}

	// If PAN data is found, wrap to JSON and return
	message = strings.TrimSuffix(message, "\n")
	response, err := json.Marshal(Message{Msg: message})
	if err != nil {
		return "", err
	}
	return string(response) + "\n", nil
}

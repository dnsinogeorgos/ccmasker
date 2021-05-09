package main

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	"github.com/theplant/luhn"
)

const maxMatches = 3

// ProcessMessage filters the message through regexp filters and returns appropriate response for rsyslog
func ProcessMessage(message string, filters []FilterGroup, numFilter *regexp.Regexp) (string, error) {
	validated := false

	for _, group := range filters {
		// If variable length pattern matches move on
		if group.Variable.MatchString(message) {
			for _, fixedPattern := range group.Fixed {
				// If fixed length pattern matches move on
				if fixedPattern.MatchString(message) {
					matchStrings := fixedPattern.FindAllString(message, maxMatches)
					for _, match := range matchStrings {
						cleanMatch := numFilter.ReplaceAllString(match, "")
						cleanInt, err := strconv.Atoi(cleanMatch)
						if err != nil {
							return "", err
						}
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
	// Otherwise wrap to JSON and save
	if validated == false {
		return "{}\n", nil
	}

	message = strings.TrimSuffix(message, "\n")
	response, err := json.Marshal(struct {
		Msg string `json:"msg"`
	}{Msg: message})
	if err != nil {
		return "", err
	}
	return string(response) + "\n", nil
}

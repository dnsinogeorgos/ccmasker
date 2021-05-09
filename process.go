package main

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/theplant/luhn"
)

const maxMatches = 3

// ProcessMessage filters the message through regexp filters and returns appropriate response for rsyslog
func ProcessMessage(message string, filters []FilterGroup, numFilter *regexp.Regexp) string {
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
							log.Fatalf("error: could not convert Luhn verified numFiltered value to string")
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
		return "{}\n"
	}

	message = strings.TrimSuffix(message, "\n")
	response, err := json.Marshal(struct {
		Msg string `json:"msg"`
	}{Msg: message})
	if err != nil {
		printErrorWithErrorHandling("Error %s occured during json Marshal of %s\n", err, message)
	}
	return string(response) + "\n"
}

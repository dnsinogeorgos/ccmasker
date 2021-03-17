package main

import (
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/moovweb/rubex"
)

// Get pointers to values to minimize copying
// If PAN data is found mask PAN in place
func processMessage(message string, filters map[string]*rubex.Regexp) string {
	var matches bool

	for mask, filter := range filters {
		if filter.MatchString(message) {
			if matches == false {
				matches = true
			}
			message = filter.ReplaceAllString(message, mask)
		}
	}

	// If PAN data isn't found, return empty JSON
	// Otherwise wrap to JSON and save
	if matches == false {
		return "{}\n"
	} else {
		var err error
		message = strings.TrimSuffix(message, "\n")
		response, err := jsoniter.Marshal(struct {
			Msg string `json:"msg"`
		}{Msg: message})
		if err != nil {
			printErrorWithErrorHandling("Error %s occured during json Marshal of %s\n", err, message)
		}
		return string(response) + "\n"
	}
}

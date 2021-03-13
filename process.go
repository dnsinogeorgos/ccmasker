package main

import (
	"encoding/json"
	"regexp"
)

// Get pointers to values to minimize copying
// If PAN data is found mask PAN in place
func processMessage(matches *bool, message *string, response *[]byte, filters map[string]*regexp.Regexp) {
	*matches = false
	for mask, filter := range filters {
		if filter.MatchString(*message) {
			if *matches == false {
				*matches = true
			}
			*message = filter.ReplaceAllLiteralString(*message, mask)
		}
	}

	// If PAN data isn't found, return empty JSON
	// Otherwise wrap to JSON and save
	if *matches == false {
		*message = "{}"
	} else {
		var err error
		*response, err = json.Marshal(struct {
			Msg string `json:"msg"`
		}{Msg: *message})
		if err != nil {
			printError("Error %s occured during json Marshal of %s\n", err, *message)
		}
		*message = string(*response)
	}
}

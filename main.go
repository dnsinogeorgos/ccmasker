package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Compile and return pointer to map of filters
func compileFilters() *map[string]*regexp.Regexp {
	var s = " +\\-_" // String of characters that will be used as separators.
	var patterns = map[string]string{
		"XXXX-VISA-XXXX":       "4[0-9]{3}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?([0-9]{4}|[0-9]{1})",
		"XXXX-Master5xxx-XXXX": "5[1-5]{1}[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}",
		"XXXX-Maestro-XXXX":    "(5018|5020|5038|5893|6304|6759|6761|6762|6763)[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{0,4}[" + s + "]?[0-9]{0,3}",
		"XXXX-MaestroUK-XXXX":  "(6767[" + s + "]?70[0-9]{2}|6767[" + s + "]?74[0-9]{2})[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{0,4}[" + s + "]?[0-9]{0,3}",
		"XXXX-Master2xxx-XXXX": "2[2-7]{1}[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}",
		"XXXX-AmEx-XXXX":       "(34|37)[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{3}",
		"XXXX-DinersInt-XXXX":  "36[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{2,4}[" + s + "]?[0-9]{0,3}",
		"XXXX-DinersUSC-XXXX":  "54[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}",
	}
	filters := make(map[string]*regexp.Regexp)
	for mask, pattern := range patterns {
		filter := regexp.MustCompile(pattern)
		filters[mask] = filter
	}
	return &filters
}

type Message struct {
	Msg string `json:"msg"`
}

// If PAN data is found, mask PAN and return string.
// Otherwise return empty string.
func processMessage(text *string, filters *map[string]*regexp.Regexp) string {
	matched := false
	message := ""
	for mask, filter := range *filters {
		if filter.MatchString(*text) {
			if matched == false {
				matched = true
				message = *text
			}
			message = filter.ReplaceAllLiteralString(message, mask)
		}
	}

	if matched == false {
		return "{}"
	} else {
		jsonMessage, err := json.Marshal(Message{Msg: message})
		if err != nil {
			fmt.Printf("Error %s occured during json Marshal of %s", err, message)
		}
		return fmt.Sprintf(string(jsonMessage))
	}
}

// Open Stdin, compile regex, loop over lines
func main() {
	reader := bufio.NewReader(os.Stdin)
	filters := compileFilters()
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				fmt.Printf("Error %s occured during reader ReadString of %s\n", err, text)
			}
		}
		text = strings.TrimSuffix(text, "\n")
		fmt.Println(processMessage(&text, filters))
	}
}

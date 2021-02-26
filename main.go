package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// If PAN data is found, mask PAN and return string.
// Otherwise return empty string.
func maskPAN(text string, filters *map[string]*regexp.Regexp) string {
	for mask, filter := range *filters {
		hasPAN := filter.MatchString(text)
		if hasPAN {
			return filter.ReplaceAllLiteralString(text, mask)
		}
	}
	return ""
}

// Return message in a JSON key named "msg".
func embedToJson(text string) string {
	type Message struct {
		Msg string `json:"msg"`
	}

	message := Message{
		Msg: text,
	}
	var jsonData []byte
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error %s occured during json Marshal of %s", err, message)
	}
	return fmt.Sprintf(string(jsonData))
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
				fmt.Println("Error %s occured during reader ReadString of %s", err, text)
			}
		}
		text = strings.TrimSuffix(text, "\n")
		text = maskPAN(text, filters)

		if text == "" {
			fmt.Println("{}")
		} else {
			fmt.Println(embedToJson(text))
		}
	}
}

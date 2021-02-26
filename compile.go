package main

import (
	"fmt"
	"regexp"
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
		filter, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Println("Error %s occured during regexp Compile of %s", err, pattern)
		}
		filters[mask] = filter
	}
	return &filters
}

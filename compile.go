package main

import "regexp"

var patterns = map[string]string{
	"XXXX-VISA-XXXX":       "4[0-9]{3}[ +\\-_]?[0-9]{4}[ +\\-_]?[0-9]{4}[ +\\-_]?([0-9]{4}|[0-9]{1})",
	"XXXX-Master5xxx-XXXX": "5[1-5]{1}[0-9]{2}[ +\\-_]?[0-9]{4}[ +\\-_]?[0-9]{4}[ +\\-_]?[0-9]{4}",
	"XXXX-Maestro-XXXX":    "(5018|5020|5038|5893|6304|6759|6761|6762|6763)[ +\\-_]?[0-9]{4}[ +\\-_]?[0-9]{4}[ +\\-_]?[0-9]{0,4}[ +\\-_]?[0-9]{0,3}",
	"XXXX-MaestroUK-XXXX":  "(6767[ +\\-_]?70[0-9]{2}|6767[ +\\-_]?74[0-9]{2})[ +\\-_]?[0-9]{4}[ +\\-_]?[0-9]{0,4}[ +\\-_]?[0-9]{0,3}",
	"XXXX-Master2xxx-XXXX": "2[2-7]{1}[0-9]{2}[ +\\-_]?[0-9]{4}[ +\\-_]?[0-9]{4}[ +\\-_]?[0-9]{4}",
	"XXXX-AmEx-XXXX":       "(34|37)[0-9]{2}[ +\\-_]?[0-9]{4}[ +\\-_]?[0-9]{4}[ +\\-_]?[0-9]{3}",
	"XXXX-DinersInt-XXXX":  "36[0-9]{2}[ +\\-_]?[0-9]{4}[ +\\-_]?[0-9]{4}[ +\\-_]?[0-9]{2,4}[ +\\-_]?[0-9]{0,3}",
	"XXXX-DinersUSC-XXXX":  "54[0-9]{2}[ +\\-_]?[0-9]{4}[ +\\-_]?[0-9]{4}[ +\\-_]?[0-9]{4}",
}

// Compile and return pointer to map of filters
func compileFilters(patterns map[string]string) map[string]*regexp.Regexp {
	filters := make(map[string]*regexp.Regexp)
	for mask, pattern := range patterns {
		filter := regexp.MustCompile(pattern)
		filters[mask] = filter
	}

	return filters
}

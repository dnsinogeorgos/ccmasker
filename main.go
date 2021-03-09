package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"runtime/pprof"
	"strings"
)

// Compile and return pointer to map of filters
func compileFilters() map[string]*regexp.Regexp {
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
	return filters
}

type Message struct {
	Msg string `json:"msg"`
}

// If PAN data is found, mask PAN and return string.
// Otherwise return empty string.
func processMessage(matches *bool, text *string, message *[]byte, filters map[string]*regexp.Regexp) {
	*matches = false
	for mask, filter := range filters {
		if filter.MatchString(*text) {
			if *matches == false {
				*matches = true
			}
			*text = filter.ReplaceAllLiteralString(*text, mask)
		}
	}

	if *matches == false {
		*text = "{}"
	} else {
		var err error
		*message, err = json.Marshal(Message{Msg: *text})
		if err != nil {
			fmt.Printf("Error %s occured during json Marshal of %s", err, *text)
		}
		*text = string(*message)
	}
}

func MustClose(file *os.File, message string) {
	if err := file.Close(); err != nil {
		log.Fatal(message, err)
	}
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

// Open Stdin, compile regex, loop over lines
func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer MustClose(f, "could not close cpu profile file: ")
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	var err error = nil
	reader := bufio.NewReader(os.Stdin)
	filters := compileFilters()
	matches := false
	text := ""
	message := make([]byte, 4096)
	for {
		text, err = reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				fmt.Printf("Error %s occured during reader ReadString of %s\n", err, text)
			}
		}
		text = strings.TrimSuffix(text, "\n")
		processMessage(&matches, &text, &message, filters)
		fmt.Println(text)
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer MustClose(f, "could not close cpu profile file: ")
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

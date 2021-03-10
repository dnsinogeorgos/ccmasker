package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
)

// Function to defer closing of profile files with error handling
func mustClose(file *os.File, message string) {
	if err := file.Close(); err != nil {
		log.Fatalf(message, err)
	}
}

// Function to print to Stderr with error handling
func printError(f string, v ...interface{}) {
	_, err := fmt.Fprintf(os.Stderr, f, v...)
	if err != nil {
		log.Fatalf("Could not print to Stderr: %s\n", err)
	}
}

// Initialize flags globally
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()
	// Conditional CPU profiling
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatalf("Could not create CPU profile: %s\n", err)
		}
		defer mustClose(f, "Could not close cpu profile file: %s\n")

		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatalf("Could not start CPU profile: %s\n", err)
		}
		defer pprof.StopCPUProfile()
	}

	// Initialize variables
	var err error
	var matches bool
	var message string
	response := make([]byte, 8192)

	// Initialize reader and compile regexp filters
	reader := bufio.NewReader(os.Stdin)
	filters := compileFilters()

	for {
		// Get next message and strip trailing newline
		message, err = reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				printError("Error %s occured during reader ReadString of %s\n", err, message)
				continue
			}
		}
		message = strings.TrimSuffix(message, "\n")

		// Process message and print
		processMessage(&matches, &message, &response, filters)
		fmt.Println(message)
	}

	// Conditional memory profiling
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer mustClose(f, "could not close cpu profile file: ")

		runtime.GC()
		err = pprof.WriteHeapProfile(f)
		if err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

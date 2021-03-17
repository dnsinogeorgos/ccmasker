package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/moovweb/rubex"
)

// Function to defer closing of profile files with error handling
func closeWithErrorHandling(file *os.File, message string) {
	if err := file.Close(); err != nil {
		log.Printf(message, err)
	}
}

// Function to print to Stderr with error handling
func printErrorWithErrorHandling(f string, v ...interface{}) {
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
		defer closeWithErrorHandling(f, "Could not close cpu profile file: %s\n")

		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatalf("Could not start CPU profile: %s\n", err)
		}
		defer pprof.StopCPUProfile()
	}

	// Initialize buffered reader and unbuffered writer
	reader := bufio.NewReader(os.Stdin)
	writer := io.StringWriter(os.Stdout)

	// Initialize and compile filters
	filters := make(map[string]*rubex.Regexp)
	compileFilters(filters)

	for {
		// Get next message and strip trailing newline
		message, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				printErrorWithErrorHandling("Error %s occured during reader ReadString of %s\n", err, message)
				continue
			}
		}

		// Process message and print
		response := processMessage(message, filters)
		_, err = writer.WriteString(response)
		if err != nil {
			log.Fatalf("Error %s occured during writing to Stdout\n", err)
		}
	}

	// Conditional memory profiling
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatalf("could not create memory profile: %s\n", err)
		}
		defer closeWithErrorHandling(f, "could not close memory profile file: %s\n")

		runtime.GC()
		err = pprof.WriteHeapProfile(f)
		if err != nil {
			log.Fatalf("could not write memory profile: %s\n", err)
		}
	}
}

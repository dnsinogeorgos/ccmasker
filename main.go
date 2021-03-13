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
)

// Function to defer closing of profile files with error handling
func closeWithErrorHandling(file *os.File, message string) {
	if err := file.Close(); err != nil {
		log.Printf(message, err)
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

	// Initialize variables
	var err error
	var matches bool
	var message string
	response := make([]byte, 0)

	// Conditional CPU profiling
	if *cpuprofile != "" {
		var fc *os.File
		fc, err = os.Create(*cpuprofile)
		if err != nil {
			log.Fatalf("Could not create CPU profile: %s\n", err)
		}
		defer closeWithErrorHandling(fc, "Could not close cpu profile file: %s\n")

		err = pprof.StartCPUProfile(fc)
		if err != nil {
			log.Fatalf("Could not start CPU profile: %s\n", err)
		}
		defer pprof.StopCPUProfile()
	}

	// Initialize buffered reader and unbuffered writer
	reader := bufio.NewReader(os.Stdin)
	writer := io.StringWriter(os.Stdout)

	// Compile regexp filters
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

		// Process message and print
		processMessage(&matches, &message, &response, filters)
		_, err = writer.WriteString(message)
		if err != nil {
			log.Fatalf("Error %s occured during writing to Stdout\n", err)
		}
	}

	// Conditional memory profiling
	if *memprofile != "" {
		var fm *os.File
		fm, err = os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer closeWithErrorHandling(fm, "could not close memory profile file: ")

		runtime.GC()
		err = pprof.WriteHeapProfile(fm)
		if err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

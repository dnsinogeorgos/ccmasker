package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"regexp"
	"runtime"
	"runtime/pprof"
)

// Function to defer closing of profile files with error handling
func closeWithErrorHandling(file *os.File, message string) {
	if err := file.Close(); err != nil {
		log.Printf(message, err)
	}
}

// Initialize flags globally
var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile = flag.String("memprofile", "", "write memory profile to `file`")
)

func main() {
	flag.Parse()
	// Conditional CPU profiling
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatalf("could not create CPU profile: %s\n", err)
		}
		defer closeWithErrorHandling(f, "could not close cpu profile file: %s\n")

		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatalf("could not start CPU profile: %s\n", err)
		}
		defer pprof.StopCPUProfile()
	}

	// Initialize buffered reader and unbuffered writer
	reader := bufio.NewReader(os.Stdin)
	writer := io.StringWriter(os.Stdout)

	// Initialize and compile filters
	numFilter, err := regexp.Compile("[^0-9]")
	if err != nil {
		log.Fatalf("could not compile number filter: %s", err)
	}
	filters := CompileFilters()

	for {
		// Get next message and strip trailing newline
		message, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("could not read string: %s", err)
			}
		}

		// Process message and print
		response, err := ProcessMessage(message, filters, numFilter)
		if err != nil {
			log.Printf("could not process message: %s", err)
		}

		_, err = writer.WriteString(response)
		if err != nil {
			log.Fatalf("could not write to stdout: %s", err)
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

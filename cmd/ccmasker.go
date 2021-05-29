package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/dnsinogeorgos/ccmasker/internal/ccmasker"
)

// Function to defer closing of profile files with error handling
func closeWithErrorHandling(file *os.File, message string) {
	if err := file.Close(); err != nil {
		log.Printf(message, err)
	}
}

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

	ccmasker.Run()

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

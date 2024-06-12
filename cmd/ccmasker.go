package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

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

	printGCStats()

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

func printGCStats() {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	// Calculate average GC pause
	var totalPauseNs uint64
	for _, pause := range stats.PauseNs {
		totalPauseNs += pause
	}
	avgPauseNs := totalPauseNs / uint64(len(stats.PauseNs))

	// Calculate average objects collected per GC cycle
	avgObjectsCollected := float64(stats.NumGC) / float64(stats.NumGC)

	log.Println("Garbage Collection Stats:")
	log.Printf("  Number of GCs: %v\n", stats.NumGC)
	log.Printf("  Total GC Pause Time: %v ms\n", stats.PauseTotalNs/1e6)
	log.Printf("  Average GC Pause Time: %v ms\n", avgPauseNs/1e6)
	log.Printf("  Total Alloc: %v bytes\n", stats.TotalAlloc)
	log.Printf("  Heap Alloc: %v bytes\n", stats.HeapAlloc)
	log.Printf("  Heap Objects: %v\n", stats.HeapObjects)
	log.Printf("  Average Objects Collected per GC: %v\n", avgObjectsCollected)
	log.Printf("  Last GC Time: %v\n", time.Unix(0, int64(stats.LastGC)))
}

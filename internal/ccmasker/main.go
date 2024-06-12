package ccmasker

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
)

func Run() {
	// Initialize buffered reader and unbuffered writer
	reader := bufio.NewReader(os.Stdin)
	writer := io.Writer(os.Stdout)

	// Initialize and compile filters
	numFilter := regexp.MustCompile("[^0-9]")
	filters := compileFilters()

	// Main loop
	for {
		// Get next message and strip trailing newline
		message, err := reader.ReadSlice('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("could not read string: %s", err)
			}
		}

		// Process message and print to stdout
		err = ProcessMessage(writer, message, filters, numFilter)
		if err != nil {
			log.Printf("could not process message: %s", err)
		}
	}
}

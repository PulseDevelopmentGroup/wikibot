package main

import (
	"fmt"
	"flag"
	"os"
	"log"
)

func main() {
	/* Flags */
	var (
		start = flag.String(
			"start",
			"https://en.wikipedia.org/wiki/Go_(programming_language)",
			"Wikipedia page `URL` to start at (starting with \"https://\")",
		)
		end = flag.String(
			"end",
			"https://en.wikipedia.org/wiki/Google",
			"Wikipedia page `URL` to end at (starting with \"https://\")",
		)
		maxThreads = flag.Int(
			"threads",
			10,
			"Maximum `number` of threads to use when searching",
		)
	)
	flag.Usage = func() {
		fmt.Fprintf(
			os.Stderr,
			`Usage: %v [options] -start [Wikipedia Page] -end [Wikipedia Page]

Finds a link between two Wikipedia pages and prints it.

Options:
`,
			os.Args[0],
		)
		flag.PrintDefaults()
	}
	flag.Parse()

	fmt.Printf("Start: %q, End: %q, Max Threads: %v\n", *start, *end, *maxThreads)
}
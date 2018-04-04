// This is a tool that helps either *extract* or *filter out* lines of text
// between a START marker and an END marker.
// The START and END can be regexes.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jvoung/gobetween"
)

var (
	startRe = flag.String("start", "", "Regex for START delimiter (required)")
	endRe   = flag.String("end", "", "Regex for END delimiter (required)")
	invert  = flag.Bool("v", false, "Invert matches. Instead of printing "+
		"everything between START and END, print lines that are not between "+
		"START and END)")
)

func main() {
	flag.Parse()
	if *startRe == "" || *endRe == "" {
		fmt.Println("must specify a start/end regex")
		flag.PrintDefaults()
		os.Exit(1)
	}
	remaining := flag.Args()
	var input *os.File
	if len(remaining) == 0 {
		input = os.Stdin
	} else if len(remaining) == 1 {
		var err error
		input, err = os.Open(remaining[0])
		if err != nil {
			panic(err)
		}
		defer input.Close()
	} else {
		fmt.Println("Can specify at most one input file")
		os.Exit(1)
	}
	// TODO(jvoung): pass a scanner instead?
	gobetween.PrintBetween(input, *startRe, *endRe)
}

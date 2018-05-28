// This is a tool that helps prints all lines between START markers and
// an END markers. The matching can also be inverted to print all lines
// that are *not* between the markers.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/jvoung/grepbetween"
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
		fmt.Printf("must specify a start/end regex (given start=%s, end=%s)",
			*startRe, *endRe)
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
		fmt.Printf(
			"Can specify at most one input as argument, given %d arguments",
			len(remaining))
		os.Exit(1)
	}
	inScan := bufio.NewScanner(input)
	grepbetween.PrintBetween(inScan, os.Stdout, *startRe, *endRe, *invert)
}

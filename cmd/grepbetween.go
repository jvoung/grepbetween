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

type stringArray []string

func (i *stringArray) String() string {
	return fmt.Sprint(*i)
}
func (i *stringArray) Set(s string) error {
	*i = append(*i, s)
	return nil
}

var (
	startRe stringArray
	endRe   stringArray
	invert  = flag.Bool("v", false, "Invert matches. Instead of printing "+
		"everything between START and END, print lines that are not between "+
		"START and END)")
)

func init() {
	flag.Var(&startRe, "start", "Regex for START delimiter (required)")
	flag.Var(&endRe, "end", "Regex for END delimiter (required)")
}

func main() {
	flag.Parse()
	if len(startRe) == 0 || len(endRe) == 0 {
		fmt.Printf("must specify a start/end regex (given start=%s, end=%s)\n",
			startRe, endRe)
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
	grepbetween.PrintBetween(inScan, os.Stdout, startRe, endRe, *invert)
}

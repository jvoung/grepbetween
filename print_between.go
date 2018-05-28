package grepbetween

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

// PrintBetween prints all lines that are between startRe and endRe pairs.
// If invert is true, then prints the inverse (lines not between startRe
// and endRe).
func PrintBetween(input *bufio.Scanner, output io.Writer,
	startRe string, endRe string, invert bool) {
	// TODO(jvoung): handle a list of start/end pairs.
	inMatch := false
	if invert {
		inMatch = true
		startRe, endRe = endRe, startRe
	}
	for input.Scan() {
		line := input.Text()
		if inMatch {
			matched, err := regexp.MatchString(endRe, line)
			if err != nil {
				panic(err)
			}
			if matched {
				inMatch = !inMatch
				continue
			}
			fmt.Fprintln(output, line)
		} else {
			matched, err := regexp.MatchString(startRe, line)
			if err != nil {
				panic(err)
			}
			if matched {
				inMatch = !inMatch
			}
		}
	}
}

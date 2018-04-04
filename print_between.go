package gobetween

import (
	"os"
	"regexp"
)

// PrintBetween prints all lines in the input that are between
// startRe and endRe.
func PrintBetween(input *os.File, startRe string, endRe string) {
	// TODO(jvoung): handle a list of start/end pairs.
	// TODO(jvoung): take a Writer to fill the output, instead
	// of just printing to stdout.
	// TODO(jvoung): take a Reader or Scanner to get input
	// instead of requiring a File.
	var line string
	regexp.MatchString(startRe, line)
}

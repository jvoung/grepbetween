package grepbetween

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

// PrintBetween prints all lines that are between any pair of patterns
// in startPatterns and endPatterns. If invert is true, then prints the
// inverse (lines not between startPatterns and endPatterns).
func PrintBetween(input *bufio.Scanner, output io.Writer,
	startPatterns []string, endPatterns []string, invert bool) {
	// TODO(jvoung): handle a list of start/end pairs.
	inMatch := false
	if invert {
		inMatch = true
		startPatterns, endPatterns = endPatterns, startPatterns
	}
	startRegexes := compilePatterns(startPatterns)
	endRegexes := compilePatterns(endPatterns)
	for input.Scan() {
		line := input.Text()
		if inMatch {
			if anyMatch(endRegexes, line) {
				inMatch = !inMatch
				continue
			}
			fmt.Fprintln(output, line)
		} else {
			if anyMatch(startRegexes, line) {
				inMatch = !inMatch
			}
		}
	}
}

func compilePatterns(patterns []string) []*regexp.Regexp {
	var result []*regexp.Regexp
	for _, p := range patterns {
		re, err := regexp.Compile(p)
		if err != nil {
			panic(err)
		}
		result = append(result, re)
	}
	return result
}

func anyMatch(regexes []*regexp.Regexp, s string) bool {
	for _, r := range regexes {
		if r.MatchString(s) {
			return true
		}
	}
	return false
}

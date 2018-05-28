package grepbetween

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestEmptyInput(t *testing.T) {
	input := inputScanner("")
	expected := ""
	startRe := "^.*START.*"
	endRe := "^.*END.*"
	assertPrinting(t, input, startRe, endRe).isEqualTo(expected)
}

func TestEmptyMatch(t *testing.T) {
	input := inputScanner("stuff\nmore stuff")
	expected := ""
	startRe := "^.*START.*"
	endRe := "^.*END.*"
	assertPrinting(t, input, startRe, endRe).isEqualTo(expected)
}

func TestMatchMost(t *testing.T) {
	input := inputScanner("===START===\nl1\nl2\nl3\n===END===")
	expected := "l1\nl2\nl3\n"
	startRe := "^.*START.*"
	endRe := "^.*END.*"
	assertPrinting(t, input, startRe, endRe).isEqualTo(expected)
}

func TestPrintBetweenOnce(t *testing.T) {
	input := inputScanner(`first line
===START_PAT===
second
third
===END_PAT===
fourth`)
	expected := `second
third
`
	startRe := "^.*START.*"
	endRe := "^.*END.*"
	assertPrinting(t, input, startRe, endRe).isEqualTo(expected)
}

func TestPrintBetweenTwice(t *testing.T) {
	input := inputScanner(`first line
===START_PAT===
second
third
===END_PAT===
fourth
===START_PAT===
fifth
===END_PAT===`)
	expected := `second
third
fifth
`
	startRe := "^.*START.*"
	endRe := "^.*END.*"
	assertPrinting(t, input, startRe, endRe).isEqualTo(expected)
}

func TestPrintBetweenOnceInverted(t *testing.T) {
	input := inputScanner(`first line
===START_PAT===
second
third
===END_PAT===
fourth`)
	expected := `first line
fourth
`
	startRe := "^.*START.*"
	endRe := "^.*END.*"
	assertPrintingInverted(t, input, startRe, endRe).isEqualTo(expected)
}

func TestMultiStartEnd(t *testing.T) {
	input := inputScanner(`first line
===START_PAT===
===START_PAT===
===START_PAT===
second
third
===END_PAT===
===END_PAT===
fourth`)
	expected := `===START_PAT===
===START_PAT===
second
third
`
	startRe := "^.*START.*"
	endRe := "^.*END.*"
	assertPrinting(t, input, startRe, endRe).isEqualTo(expected)
}

func inputScanner(s string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(s))
}

type stringSubject struct {
	t      *testing.T
	actual string
}

func assertPrinting(t *testing.T, input *bufio.Scanner,
	startRe, endRe string) stringSubject {
	output := bytes.Buffer{}
	PrintBetween(input, &output, startRe, endRe, false)
	return assertThat(t, output.String())
}

func assertPrintingInverted(t *testing.T, input *bufio.Scanner,
	startRe, endRe string) stringSubject {
	output := bytes.Buffer{}
	PrintBetween(input, &output, startRe, endRe, true)
	return assertThat(t, output.String())
}

func assertThat(t *testing.T, actual string) stringSubject {
	return stringSubject{t, actual}
}

func (subj stringSubject) isEqualTo(expected string) {
	if expected != subj.actual {
		subj.t.Errorf("Expected != actual. (expected=\"%s\", actual=\"%s\")",
			expected, subj.actual)
	}
}

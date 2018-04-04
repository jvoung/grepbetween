package gobetween

import (
	"os"
	"testing"
)

func TestPrintBetween(t *testing.T) {
	type args struct {
		input   *os.File
		startRe string
		endRe   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintBetween(tt.args.input, tt.args.startRe, tt.args.endRe)
		})
	}
}

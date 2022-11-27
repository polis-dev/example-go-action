package main

import (
	"os"

	//lint:ignore ST1001 its okay. just chill.
	. "action"

	gha "github.com/sethvargo/go-githubactions"
)

func main() {
	r, err := New()
	if err != nil {
		gha.Errorf("error initializing context: %v", err)
		os.Exit(1)
	}

	r.OutputDiagnosticsGroup()
	r.AddStepSummary()
}

package main

import (
	"os"

	"action"

	gha "github.com/sethvargo/go-githubactions"
)

func main() {
	r, err := action.New()
	if err != nil {
		gha.Errorf("error initializing context: %v", err)
		os.Exit(1)
	}

	r.OutputDiagnosticsGroup()
	r.AddStepSummary()
}

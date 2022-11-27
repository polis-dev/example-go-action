package action

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	lo "github.com/samber/lo"
	gha "github.com/sethvargo/go-githubactions"
)

// Run represents the GitHub:Run at runtime.
type Run struct {
	// We wrap githubactions.Action to add our own fields.
	*gha.Action
	// Context is the context provided by the GitHub Actions runtime.
	*gha.GitHubContext
	// action is the GitHub Actions instance.
	// Env is a map of key-value pairs from the environment.
	Env map[string]string
	// Debug represents if debug mode is enabled.
	Debug bool
	// Command is an arbitrary value provided via an "input".
	Command string
}

// New creates a new Action from the environment/inputs/etc.
func New() (r *Run, err error) {
	r = &Run{
		Action: gha.WithFieldsMap(map[string]string{
			"val": gha.GetInput("val"),
		}),
		Command: gha.GetInput("command"),
		Debug:   gha.GetInput("debug") == "true",
		Env:     EnvironMap(),
	}
	// load the context provided by github.
	if r.GitHubContext, err = gha.Context(); err != nil {
		return nil, err
	} else if r.RunID == 0 || r.EventName == "" {
		// sanity check.
		fmt.Println("i think i'm confused... is this a valid workflow runtime?")
	}
	return
}

// String returns a string representation of the Context.
func (r *Run) String() string {
	return fmt.Sprintf("%s on %s@%s (Run #%d)", r.Workflow, r.Repository, r.SHA , r.RunID)
}

// OutputDiagnosticsGroup pretty-prints a JSON representation of this Context
// for diagnostics/humans.
func (r *Run) OutputDiagnosticsGroup() {
	r.Action.Group("Context")
	b, _ := json.MarshalIndent(r, "", "  ")
	fmt.Println(string(b))
	r.Action.EndGroup()
}

func (r *Run) AddStepSummary() error {
	if r.StepSummary == "" {
		// we silently skip printing the step summary if it's not set.
		return nil
	}

	return r.Action.AddStepSummaryTemplate(`
## {{ .Heading }}

- Thanks to {{.Actor}} :moon: for triggering this Workflow!

`, map[string]string{
		"Heading": r.String(),
		"Actor":   r.Actor,
	})
}

// EnvironMap builds a map of key-value pairs from the []string returned by os.Environ().
func EnvironMap() map[string]string {
	return lo.Reduce(os.Environ(), func(agg map[string]string, item string, i int) map[string]string {
		parts := strings.SplitN(item, "=", 2)
		if len(parts) == 2 {
			agg[parts[0]] = parts[1]
		}
		return agg
	}, map[string]string{})
}

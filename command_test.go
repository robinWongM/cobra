package cobra

import (
	"bytes"
	"testing"
)

func TestSimpleCommandWithArgs(t *testing.T) {
	itDidRun := 0

	// Make a root command
	rootCmd := &Command{
		Use: "root",
		Run: func(_ *Command, _ []string) {
			itDidRun++
		},
	}

	// How do we capture the output of a Command?
	// Override it.
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.Execute()

	// Test if there exists unexpected output
	got := buf.String()
	expected := ""

	if got != expected {
		t.Errorf("Unexpected output: %v", got)
	}

	// Test if the Run function has run once
	if itDidRun != 1 {
		t.Errorf("Run function didn't run as expected: %v", itDidRun)
	}
}

package cobra

import (
	"bytes"
	"reflect"
	"testing"
)

func executeCommand(c *Command, args ...string) string {
	// How do we capture the output of a Command?
	// Override it.
	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)
	c.Execute()

	return buf.String()
}

func TestSimpleCommand(t *testing.T) {
	itDidRun := 0

	// Make a root command
	rootCmd := &Command{
		Use: "root",
		Run: func(_ *Command, _ []string) {
			itDidRun++
		},
	}

	// Test if there exists unexpected output
	got := executeCommand(rootCmd)
	expected := ""

	if got != expected {
		t.Errorf("Unexpected output: %v", got)
	}

	// Test if the Run function has run once
	if itDidRun != 1 {
		t.Errorf("Run function didn't run as expected: %v", itDidRun)
	}
}

func TestSingleCommandWithArgs(t *testing.T) {
	var rootCmdArgs []string

	rootCmd := &Command{
		Use: "root",
		Run: func(_ *Command, args []string) {
			rootCmdArgs = args
		},
	}

	got := executeCommand(rootCmd, "-p", "8080")
	expected := ""

	if got != expected {
		t.Errorf("Unexpected output: %v", got)
	}

	if !reflect.DeepEqual(rootCmdArgs, []string{"-p", "8080"}) {
		t.Errorf("Run function didn't receive expected args: %v", rootCmdArgs)
	}
}

package cobra

import (
	"bytes"
	"reflect"
	"testing"
)

func executeCommand(c *Command, args ...string) (string, error) {
	// How do we capture the output of a Command?
	// Override it.
	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)
	_, err := c.Execute()

	return buf.String(), err
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
	got, err := executeCommand(rootCmd)
	expected := ""

	if err != nil {
		t.Errorf("Unexpected error: %e", err)
	}
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

	got, err := executeCommand(rootCmd, "one", "two")
	expected := ""

	if err != nil {
		t.Errorf("Unexpected error: %e", err)
	}
	if got != expected {
		t.Errorf("Unexpected output: %v", got)
	}

	if !reflect.DeepEqual(rootCmdArgs, []string{"one", "two"}) {
		t.Errorf("Run function didn't receive expected args: %v", rootCmdArgs)
	}
}

func emptyRun(_ *Command, _ []string) {}

func TestChildCommand(t *testing.T) {
	var childCmdArgs []string
	rootCmd := &Command{Use: "root", Run: emptyRun}
	childCmd := &Command{
		Use: "child",
		Run: func(_ *Command, args []string) {
			childCmdArgs = args
		},
	}
	rootCmd.AddCommand(childCmd)

	got, err := executeCommand(rootCmd, "child", "one", "two")
	expected := ""

	if err != nil {
		t.Errorf("Unexpected error: %e", err)
	}
	if got != expected {
		t.Errorf("Unexpected output: %v", got)
	}

	if !reflect.DeepEqual(childCmdArgs, []string{"one", "two"}) {
		t.Errorf("Run function didn't receive expected args: %v", childCmdArgs)
	}
}

package cobra

import (
	flag "github.com/spf13/pflag"
)

// For compatibility with Hugo.

type PositionalArgs func(cmd *Command, args []string) error

// FParseErrWhitelist configures Flag parse errors to be ignored
type FParseErrWhitelist flag.ParseErrorsWhitelist

// ShellCompDirective is a bit map representing the different behaviors the shell
// can be instructed to have once completions have been provided.
type ShellCompDirective int

// Annotations for Bash completion.
const (
	BashCompFilenameExt     = "cobra_annotation_bash_completion_filename_extensions"
	BashCompCustom          = "cobra_annotation_bash_completion_custom"
	BashCompOneRequiredFlag = "cobra_annotation_bash_completion_one_required_flag"
	BashCompSubdirsInDir    = "cobra_annotation_bash_completion_subdirs_in_dir"
)

// GenBashCompletionFile generates bash completion file.
func (c *Command) GenBashCompletionFile(filename string) error {
	// Just for compatibility
	return nil
}

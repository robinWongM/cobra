package cobra

import (
	"fmt"
	"html/template"
	"io"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	flag "github.com/spf13/pflag"
)

// For compatibility with Hugo.

type PositionalArgs func(cmd *Command, args []string) error

// FParseErrWhitelist configures Flag parse errors to be ignored
type FParseErrWhitelist flag.ParseErrorsWhitelist

// ShellCompDirective is a bit map representing the different behaviors the shell
// can be instructed to have once completions have been provided.
type ShellCompDirective int

const (
	// ShellCompDirectiveError indicates an error occurred and completions should be ignored.
	ShellCompDirectiveError ShellCompDirective = 1 << iota

	// ShellCompDirectiveNoSpace indicates that the shell should not add a space
	// after the completion even if there is a single completion provided.
	ShellCompDirectiveNoSpace

	// ShellCompDirectiveNoFileComp indicates that the shell should not provide
	// file completion even when no completion is provided.
	// This currently does not work for zsh or bash < 4
	ShellCompDirectiveNoFileComp

	// ShellCompDirectiveFilterFileExt indicates that the provided completions
	// should be used as file extension filters.
	// For flags, using Command.MarkFlagFilename() and Command.MarkPersistentFlagFilename()
	// is a shortcut to using this directive explicitly.  The BashCompFilenameExt
	// annotation can also be used to obtain the same behavior for flags.
	ShellCompDirectiveFilterFileExt

	// ShellCompDirectiveFilterDirs indicates that only directory names should
	// be provided in file completion.  To request directory names within another
	// directory, the returned completions should specify the directory within
	// which to search.  The BashCompSubdirsInDir annotation can be used to
	// obtain the same behavior but only for flags.
	ShellCompDirectiveFilterDirs

	// ===========================================================================

	// All directives using iota should be above this one.
	// For internal use.
	shellCompDirectiveMaxValue

	// ShellCompDirectiveDefault indicates to let the shell perform its default
	// behavior after completions have been provided.
	// This one must be last to avoid messing up the iota count.
	ShellCompDirectiveDefault ShellCompDirective = 0
)

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

var templateFuncs = template.FuncMap{
	"trim":                    strings.TrimSpace,
	"trimRightSpace":          trimRightSpace,
	"trimTrailingWhitespaces": trimRightSpace,
	"appendIfNotPresent":      appendIfNotPresent,
	"rpad":                    rpad,
	"gt":                      Gt,
	"eq":                      Eq,
}

// FIXME appendIfNotPresent is unused by cobra and should be removed in a version 2. It exists only for compatibility with users of cobra.

// appendIfNotPresent will append stringToAppend to the end of s, but only if it's not yet present in s.
func appendIfNotPresent(s, stringToAppend string) string {
	if strings.Contains(s, stringToAppend) {
		return s
	}
	return s + " " + stringToAppend
}

// rpad adds padding to the right of a string.
func rpad(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(template, s)
}

func trimRightSpace(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

// FIXME Gt is unused by cobra and should be removed in a version 2. It exists only for compatibility with users of cobra.

// Gt takes two types and checks whether the first type is greater than the second. In case of types Arrays, Chans,
// Maps and Slices, Gt will compare their lengths. Ints are compared directly while strings are first parsed as
// ints and then compared.
func Gt(a interface{}, b interface{}) bool {
	var left, right int64
	av := reflect.ValueOf(a)

	switch av.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		left = int64(av.Len())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		left = av.Int()
	case reflect.String:
		left, _ = strconv.ParseInt(av.String(), 10, 64)
	}

	bv := reflect.ValueOf(b)

	switch bv.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		right = int64(bv.Len())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		right = bv.Int()
	case reflect.String:
		right, _ = strconv.ParseInt(bv.String(), 10, 64)
	}

	return left > right
}

// FIXME Eq is unused by cobra and should be removed in a version 2. It exists only for compatibility with users of cobra.

// Eq takes two types and checks whether they are equal. Supported types are int and string. Unsupported types will panic.
func Eq(a interface{}, b interface{}) bool {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	switch av.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		panic("Eq called on unsupported type")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return av.Int() == bv.Int()
	case reflect.String:
		return av.String() == bv.String()
	}
	return false
}

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) error {
	t := template.New("top")
	t.Funcs(templateFuncs)
	template.Must(t.Parse(text))
	return t.Execute(w, data)
}

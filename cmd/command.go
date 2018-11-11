package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Command define the command basic operations
type Command interface {
	Init()
	Cmd() *cobra.Command
}

// baseCommand extends the cobra.Command, and defined some default operations
// , and then all sub command should extend the baseCommand to inherit the
// default operations.
type baseCommand struct {
	cmd *cobra.Command
}

// Init for base command is empty
func (b *baseCommand) Init() {}

// Cmd return the cobra Command object
func (b *baseCommand) Cmd() *cobra.Command {
	return b.cmd
}

// ExitError implement the Error interface and add some
// custom error code nad error message
type ExitError struct {
	Code   int
	ErrMsg string
}

// Error function implement the Error interface
func (e ExitError) Error() string {
	return fmt.Sprintf("Exit Code: %d, Message: %s", e.Code, e.ErrMsg)
}

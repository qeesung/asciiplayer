package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type Command interface {
	Init()
	Cmd() *cobra.Command
}

type baseCommand struct {
	cmd *cobra.Command
}

func (b *baseCommand) Init() {}

func (b *baseCommand) Cmd() *cobra.Command {
	return b.cmd
}

type ExitError struct {
	Code   int
	ErrMsg string
}

func (e ExitError) Error() string {
	return fmt.Sprintf("Exit Code: %d, Message: %s", e.Code, e.ErrMsg)
}

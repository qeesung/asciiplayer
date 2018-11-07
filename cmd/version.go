package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type VersionCommand struct {
	baseCommand
}

func (versionCommand *VersionCommand) Init() {
	versionCommand.cmd = &cobra.Command{
		Use:   "version",
		Short: "Show the version",
		Long:  SummaryTitle,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(ShortTitle + "\n" + SummaryTitle)
			return nil
		},
		Example: versionExample(),
	}
}

func versionExample() string {
	return `$ asciiplayer version`
}

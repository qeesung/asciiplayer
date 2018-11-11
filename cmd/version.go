package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// VersionCommand is responsible show the version
type VersionCommand struct {
	baseCommand
}

// Init for VersionCommand create a version command
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

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "asciiplayer",
	Short: "asciiplayer is a command line tool to play gif and video in ASCII mode",
	Long: `asciiplayer is a library that can convert gif and video to ASCII image
and provide the cli for easy use.
>>HomePage: https://github.com/qeesung/asciiplayer
>>Author  : qeesung
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

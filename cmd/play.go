// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/qeesung/asciiplayer/pkg/player"
	"github.com/spf13/cobra"
	"os"
)

var filename string
var colored string
var ratio float64
var fixedWidth int
var fixedHeight int
var fitScreen bool
var stretchedScreen bool
var delay float64
var reversed bool

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "play the gif and video in ASCII mode",
	Long: SummaryTitle + `

play command only work in terminal, decoding the gif or video
info multi frames and convert the frames to ASCII character matrix,
finally, output the matrix to stdout at a certain frequency.`,
	Run: func(cmd *cobra.Command, args []string) {
		if filename == "" { // if filename is empty
			cmd.Help()
			os.Exit(-1)
		}

		terminalPlayer, supported := player.NewTerminalPlayer(filename)
		if !supported {
			fmt.Fprintln(os.Stderr, "Not Supported file type!")
			os.Exit(-1)
		}

		playOptions := player.DefaultPlayOptions
		terminalPlayer.Play(filename, &playOptions)
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "gif filename or video filename")
}

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
	"github.com/pkg/errors"
	"github.com/qeesung/asciiplayer/pkg/player"
	"github.com/qeesung/image2ascii/convert"
	"github.com/spf13/cobra"
	"time"
)

type PlayCommand struct {
	baseCommand
	convert.Options
	Delay float64
}

func (playCommand *PlayCommand) Init() {
	playCommand.cmd = &cobra.Command{
		Use:   "play",
		Short: "Play the gif and video in ASCII mode",
		Args:  cobra.ExactArgs(1),
		Long: SummaryTitle + `

Play command only work in terminal, decoding the gif or video
info multi frames and convert the frames to ASCII character matrix,
finally, output the matrix to stdout at a certain frequency.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return playCommand.play(args)
		},
		Example: playExample(),
	}
	playCommand.addFlags()
}

func (playCommand *PlayCommand) play(args []string) error {
	filename := args[0]
	terminalPlayer, supported := player.NewTerminalPlayer(filename)
	if !supported {
		return errors.New("not supported file type")
	}

	playOptions := player.PlayOptions{}
	playOptions.Options = playCommand.Options
	playOptions.Delay = time.Duration(playCommand.Delay*1000) * time.Millisecond
	terminalPlayer.Play(filename, &playOptions)
	return nil
}

func (playCommand *PlayCommand) addFlags() {
	flagSet := playCommand.cmd.Flags()

	flagSet.Float64VarP(&playCommand.Ratio, "ratio", "r", 1.0, "Scale ratio")
	flagSet.IntVarP(&playCommand.FixedWidth, "width", "w", -1, "Scale to fixed width")
	flagSet.IntVarP(&playCommand.FixedHeight, "height", "g", -1, "Scale to fixed height")
	flagSet.BoolVarP(&playCommand.StretchedScreen, "stretched", "t", false, "Stretch the image to fit screen")
	flagSet.BoolVarP(&playCommand.Colored, "colored", "c", true, "Play with color")
	flagSet.BoolVarP(&playCommand.Reversed, "reversed", "i", false, "Play with the ascii reversed")
	flagSet.BoolVarP(&playCommand.FitScreen, "fit", "s", true, "Play fit the screen")
	flagSet.Float64VarP(&playCommand.Delay, "delay", "d", 0.15, "Play delay duration between two frames")
}

func playExample() string {
	return `$ asciiplay play hello.gif`
}

// Package player define some player that can flush the ASCII image
// to stdout.
package player

import (
	"github.com/qeesung/asciiplayer/pkg/util"
	"github.com/qeesung/image2ascii/convert"
	"time"
)

// ClearScreen is the ascii code that can clear the screen
const ClearScreen = "\033[H\033[2J"

// PlayOptions define some options for playing
type PlayOptions struct {
	convert.Options
	Delay time.Duration
}

// DefaultPlayOptions is the default and recommend options for playing
var DefaultPlayOptions = PlayOptions{
	Options: convert.DefaultOptions,
	Delay:   time.Duration(100) * time.Millisecond,
}

// Player define the default operations for playing
type Player interface {
	Play(filename string, playOptions *PlayOptions)
}

// supportedPlayerMatchers define the support player for different file type
var supportedPlayerMatchers = []struct {
	Match       func(string) bool
	Constructor func() Player
}{
	{
		Match:       util.IsGif,
		Constructor: NewGifTerminalPlayer,
	},
	{
		Match:       util.IsSupportedImage,
		Constructor: NewImageTerminalPlayer,
	},
}

// NewTerminalPlayer is factory method to create the player base on file type
func NewTerminalPlayer(filename string) (player Player, supported bool) {
	for _, matcher := range supportedPlayerMatchers {
		if matcher.Match(filename) {
			return matcher.Constructor(), true
		}
	}
	return nil, false
}

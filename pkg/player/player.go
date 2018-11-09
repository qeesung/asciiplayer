package player

import (
	"github.com/qeesung/asciiplayer/pkg/util"
	"github.com/qeesung/image2ascii/convert"
	"time"
)

const ClearScreen = "\033[H\033[2J"

type PlayOptions struct {
	convert.Options
	Delay time.Duration
}

var DefaultPlayOptions = PlayOptions{
	Options: convert.DefaultOptions,
	Delay:   time.Duration(100) * time.Millisecond,
}

type Player interface {
	Play(filename string, playOptions *PlayOptions)
}

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

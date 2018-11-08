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

// NewTerminalPlayer is factory method to create the player base on file type
func NewTerminalPlayer(filename string) (player Player, supported bool) {
	if util.IsGif(filename) {
		return NewGifTerminalPlayer(), true
	}
	return nil, false
}

package player

import (
	"fmt"
	"github.com/qeesung/asciiplayer/pkg/decoder"
	"github.com/qeesung/image2ascii/convert"
	"log"
	"os"
	"time"
)

const clearScreen = "\033[H\033[2J"

type Player interface {
	Play(filename string)
}

type TerminalPlayer struct {
	decoder   decoder.Decoder
	converter *convert.ImageConverter
}

func NewGifTerminalPlayer() Player {
	return &TerminalPlayer{
		decoder:   decoder.NewGifDeCoder(),
		converter: convert.NewImageConverter(),
	}
}

func (terminalPlayer *TerminalPlayer) Play(filename string) {
	// decode the file first
	frames, err := terminalPlayer.decoder.DecodeFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	convertOptions := convert.DefaultOptions

	stdout := os.Stdout
	for {
		for _, frame := range frames {
			asciiImageStr := terminalPlayer.converter.Image2ASCIIString(frame, &convertOptions)
			fmt.Fprint(stdout, asciiImageStr)
			time.Sleep(time.Duration(100) * time.Millisecond)
			fmt.Fprint(stdout, clearScreen)
		}
	}
}

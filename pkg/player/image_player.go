package player

import (
	"fmt"
	"github.com/qeesung/asciiplayer/pkg/decoder"
	"github.com/qeesung/image2ascii/convert"
	"log"
	"os"
)

// ImageTerminalPlayer responsible for playing the image in the terminal
type ImageTerminalPlayer struct {
	decoder   decoder.Decoder
	converter *convert.ImageConverter
}

// NewImageTerminalPlayer create a new ImageTerminalPlayer object
func NewImageTerminalPlayer() Player {
	return &ImageTerminalPlayer{
		decoder:   decoder.NewImageDecoder(),
		converter: convert.NewImageConverter(),
	}
}

// Play for ImageTerminalPlayer flush the image to the stdout
func (player *ImageTerminalPlayer) Play(filename string, playOptions *PlayOptions) {
	// decode the file first
	frames, err := player.decoder.DecodeFromFile(filename, nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(frames) == 0 {
		log.Fatal("missing frames")
	}
	frame := frames[0]

	convertOptions := playOptions.Options

	asciiImageStr := player.converter.Image2ASCIIString(frame, &convertOptions)
	fmt.Fprint(os.Stdout, asciiImageStr)
}

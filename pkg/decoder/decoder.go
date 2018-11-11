// Package decoder is responsible for split the video or gif to frames
package decoder

import (
	"github.com/qeesung/asciiplayer/pkg/util"
	"image"
	"io"
)

// Decoder interface define the basic operation to decode the gif or video
type Decoder interface {
	// Decode decode a file into multi frames
	Decode(r io.Reader, progress chan<- int) (frames []image.Image, err error)
	DecodeFromFile(filename string, progress chan<- int) (frames []image.Image, err error)
}

var supportedDecoderMatchers = []struct {
	Match       func(string) bool
	Constructor func() Decoder
}{
	{
		Match:       util.IsGif,
		Constructor: NewGifDeCoder,
	},
	{
		Match:       util.IsSupportedImage,
		Constructor: NewImageDecoder,
	},
}

// NewDecoder is factory method to create the player base on file type
func NewDecoder(filename string) (decoder Decoder, supported bool) {
	for _, matcher := range supportedDecoderMatchers {
		if matcher.Match(filename) {
			return matcher.Constructor(), true
		}
	}
	return nil, false
}

// Package encoder responsible for merging multi frames to a gif or video
package encoder

import (
	"github.com/qeesung/asciiplayer/pkg/util"
	"github.com/qeesung/image2ascii/convert"
	"image"
	"io"
	"time"
)

// Encoder interface is used to encode the multi frames to a gif file
// or encode frames to a video
type Encoder interface {
	Encode(writer io.Writer, frames []image.Image, progress chan<- int) error
	EncodeToFile(filename string, frames []image.Image, progress chan<- int) error
}

// EncodeOptions define the required options to encode
type EncodeOptions struct {
	convert.Options
	Delay time.Duration
}

// DefaultEncodeOptions is default and recommend options for encoding
var DefaultEncodeOptions = EncodeOptions{
	Options: convert.DefaultOptions,
	Delay:   time.Duration(100) * time.Millisecond,
}

// supportedEncoderMatcher define the supported encoders for different file type.
var supportedEncoderMatcher = []struct {
	Match       func(string) bool
	Constructor func() Encoder
}{
	{
		Match:       util.IsGif,
		Constructor: NewGifEncoder,
	},
	{
		Match:       util.IsSupportedImage,
		Constructor: NewImageEncoder,
	},
}

// NewEncoder is a factory method to create a new encoder by file type
func NewEncoder(filename string) (encoder Encoder, supported bool) {
	for _, matcher := range supportedEncoderMatcher {
		if matcher.Match(filename) {
			return matcher.Constructor(), true
		}
	}
	return nil, false
}

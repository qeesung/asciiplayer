// encoder package responsible for merging multi frames to a gif or video
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

type EncodeOptions struct {
	convert.Options
	Delay time.Duration
}

var DefaultEncodeOptions = EncodeOptions{
	Options: convert.DefaultOptions,
	Delay:   time.Duration(100) * time.Millisecond,
}

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

func NewEncoder(filename string) (encoder Encoder, supported bool) {
	for _, matcher := range supportedEncoderMatcher {
		if matcher.Match(filename) {
			return matcher.Constructor(), true
		}
	}
	return nil, false
}

// decoder package is responsible for split the video or gif to frames
package decoder

import (
	"image"
	"io"
)

type Decoder interface {
	// Decode decode a file into multi frames
	Decode(r io.Reader) (frames []image.Image, err error)
	DecodeFromFile(filename string) (frames []image.Image, err error)
}

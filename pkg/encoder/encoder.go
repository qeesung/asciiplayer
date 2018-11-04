// encoder package responsible for merging multi frames to a gif or video
package encoder

import (
	"image"
	"io"
)

// Encoder interface is used to encode the multi frames to a gif file
// or encode frames to a video
type Encoder interface {
	Encode(writer io.Writer, frames []image.Image) error
	EncodeToFile(filename string, frames []image.Image) error
}

// encoder package responsible for merging multi frames to a gif or video
package encoder

import (
	"image"
	"io"
)

type Encoder interface {
	Encode(writer io.Writer, frames []image.Image) error
	EncodeToFile(filename string, frames []image.Image) error
}

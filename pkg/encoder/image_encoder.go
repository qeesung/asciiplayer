package encoder

import (
	"github.com/pkg/errors"
	"github.com/qeesung/asciiplayer/pkg/util"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

// ImageEncoder is responsible for encoding image
type ImageEncoder struct {
	Filename string
}

// NewImageEncoder create a new ImageEncoder
func NewImageEncoder() Encoder {
	return &ImageEncoder{}
}

// Encode for ImageEncoder just encode the image, if the frame slice contain many frame, only
// pick up the first one.
func (encoder *ImageEncoder) Encode(writer io.Writer, frames []image.Image, progress chan<- int) error {
	if progress != nil {
		defer close(progress)
	}
	// only encode the first frame
	if len(frames) == 0 {
		return errors.New("missing frame")
	}
	frame := frames[0]

	if progress != nil {
		progress <- 1
	}

	if util.IsJPG(encoder.Filename) {
		return jpeg.Encode(writer, frame, nil)
	}
	return png.Encode(writer, frame)
}

// EncodeToFile encode the frames to the file by filename
func (encoder *ImageEncoder) EncodeToFile(filename string, frames []image.Image, progress chan<- int) error {
	encoder.Filename = filename
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return encoder.Encode(f, frames, progress)
}

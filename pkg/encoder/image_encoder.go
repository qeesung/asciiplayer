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

type ImageEncoder struct {
	Filename string
}

func NewImageEncoder() Encoder {
	return &ImageEncoder{}
}

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
	} else {
		return png.Encode(writer, frame)
	}
}

func (encoder *ImageEncoder) EncodeToFile(filename string, frames []image.Image, progress chan<- int) error {
	encoder.Filename = filename
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return encoder.Encode(f, frames, progress)
}

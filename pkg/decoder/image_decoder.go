package decoder

import (
	"image"
	"io"
	"os"
)

func NewImageDecoder() Decoder {
	return &ImageDecoder{}
}

type ImageDecoder struct {
}

func (decoder *ImageDecoder) Decode(reader io.Reader, progress chan<- int) (frames []image.Image, err error) {
	if progress != nil {
		defer close(progress)
	}

	img, _, err := image.Decode(reader)

	if progress != nil {
		progress <- 1
	}
	if err != nil {
		return nil, err
	}

	return []image.Image{img}, nil
}

func (decoder *ImageDecoder) DecodeFromFile(filename string, progress chan<- int) (frames []image.Image, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return decoder.Decode(f, progress)
}

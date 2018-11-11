package decoder

import (
	"image"
	"io"
	"os"
)

// NewImageDecoder create a new ImageDecoder
func NewImageDecoder() Decoder {
	return &ImageDecoder{}
}

// ImageDecoder is responsible for decoding image
type ImageDecoder struct {
}

// Decode for ImageDecoder decoding a image and return a frame slice that only
// contain one frame.
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

// DecodeFromFile decode the file to frames by filename
func (decoder *ImageDecoder) DecodeFromFile(filename string, progress chan<- int) (frames []image.Image, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return decoder.Decode(f, progress)
}

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

func (decoder *ImageDecoder) Decode(reader io.Reader) (frames []image.Image, err error) {
	img, _, err := image.Decode(reader)

	if err != nil {
		return nil, err
	}

	return []image.Image{img}, nil
}

func (decoder *ImageDecoder) DecodeFromFile(filename string) (frames []image.Image, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return decoder.Decode(f)
}

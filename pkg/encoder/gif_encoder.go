package encoder

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"io"
	"os"
)

// GifEncoder is responsible for encode the frames into gif
type GifEncoder struct {
}

// NewGifEncoder create a new encoder
func NewGifEncoder() Encoder {
	return &GifEncoder{}
}

// Encode encode frames into a io writer
func (gifEncoder *GifEncoder) Encode(writer io.Writer, frames []image.Image) error {
	palette := append(palette.WebSafe, color.Transparent)
	outGif := &gif.GIF{}
	for _, frame := range frames {
		bounds := frame.Bounds()
		paletteImage := image.NewPaletted(bounds, palette)
		draw.Draw(paletteImage, bounds, frame, image.ZP, draw.Src)
		outGif.Image = append(outGif.Image, paletteImage)
		outGif.Delay = append(outGif.Delay, 1)
	}
	return gif.EncodeAll(writer, outGif)
}

// EncodeToFile encode frames into a file by file name
func (gifEncoder *GifEncoder) EncodeToFile(gifFilename string, frames []image.Image) error {
	f, err := os.Create(gifFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	return gifEncoder.Encode(f, frames)
}

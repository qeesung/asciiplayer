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

type GifEncoder struct {
}

func NewGifEncoder() Encoder {
	return &GifEncoder{}
}

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

func (gifEncoder *GifEncoder) EncodeToFile(gifFilename string, frames []image.Image) error {
	f, err := os.Create(gifFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	return gifEncoder.Encode(f, frames)
}

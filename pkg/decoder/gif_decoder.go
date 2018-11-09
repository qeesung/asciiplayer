package decoder

import (
	"image"
	"image/draw"
	"image/gif"
	"io"
	"os"
)

// GifDecoder responsible for decoding the gif and implement the Decoder interface
type GifDecoder struct {
}

// NewGifDeCoder create a new gif decoder
func NewGifDeCoder() Decoder {
	return &GifDecoder{}
}

// Decode for GifDeCoder decode the gif file to multi frames
func (gifDecoder *GifDecoder) Decode(reader io.Reader, progress chan<- int) (frames []image.Image, err error) {
	if progress != nil {
		defer close(progress)
	}
	gifImage, err := gif.DecodeAll(reader)

	if err != nil {
		return nil, err
	}

	imgWidth, imgHeight := gifDecoder.getGifDimensions(gifImage)

	overPaintImage := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	draw.Draw(overPaintImage, overPaintImage.Bounds(), gifImage.Image[0], image.ZP, draw.Src)

	for _, srcImg := range gifImage.Image {
		draw.Draw(overPaintImage, overPaintImage.Bounds(), srcImg, image.ZP, draw.Over)
		frame := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
		draw.Draw(frame, frame.Bounds(), overPaintImage, image.ZP, draw.Over)
		frames = append(frames, frame)
		if progress != nil {
			progress <- 1
		}
	}

	return frames, nil
}

// DecodeFromFile decode the gif file by filename to multi frames
func (gifDecoder *GifDecoder) DecodeFromFile(gifFilename string, progress chan<- int) (frames []image.Image, err error) {
	f, err := os.Open(gifFilename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return gifDecoder.Decode(f, progress)
}

// getGifDimensions get the gif dimension
func (gifDecoder *GifDecoder) getGifDimensions(gif *gif.GIF) (x, y int) {
	var lowestX int
	var lowestY int
	var highestX int
	var highestY int

	for _, img := range gif.Image {
		if img.Rect.Min.X < lowestX {
			lowestX = img.Rect.Min.X
		}
		if img.Rect.Min.Y < lowestY {
			lowestY = img.Rect.Min.Y
		}
		if img.Rect.Max.X > highestX {
			highestX = img.Rect.Max.X
		}
		if img.Rect.Max.Y > highestY {
			highestY = img.Rect.Max.Y
		}
	}

	return highestX - lowestX, highestY - lowestY
}

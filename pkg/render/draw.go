// render package convert the char pixel matrix that converted from
// the image2ascii to a png image that draw all ascii chars to the image
package render

import (
	"errors"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/qeesung/image2ascii/ascii"
	"github.com/qeesung/image2ascii/convert"
	"golang.org/x/image/font/gofont/gomono"
	"image"
	"image/color"
	"reflect"
)

// DrawOptions control how to draw the pixel to the image
type DrawOptions struct {
	TTF             []byte  // draw font
	FontSize        float64 // font size
	Colored         bool
	BackGroundColor color.RGBA
	ForeGroundColor color.RGBA
}

// DefaultDrawOptions is default draw options
var DefaultDrawOptions = DrawOptions{
	TTF:      gomono.TTF,
	FontSize: 20,
	Colored:  true,
	BackGroundColor: color.RGBA{ // default background color is black
		R: 0,
		G: 0,
		B: 0,
		A: 255,
	},
	ForeGroundColor: color.RGBA{ // default foreground color is white
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	},
}

// Drawer interface define the operation that draw char pixel to image
type Drawer interface {
	DrawCharPixelMatrix2Image(charPixelMatrix [][]ascii.CharPixel, options DrawOptions) (img image.Image, err error)
	BatchConvertThenDraw(frames []image.Image, convertOptions convert.Options, drawOptions DrawOptions, progress chan<- int) (asciiImages []image.Image, err error)
}

// ImageDrawer implement the drawer interface
type ImageDrawer struct {
}

// NewImageDrawer create a new image drawer
func NewImageDrawer() Drawer {
	return &ImageDrawer{}
}

func (drawer *ImageDrawer) BatchConvertThenDraw(frames []image.Image,
	convertOptions convert.Options, drawOptions DrawOptions, progress chan<- int) (asciiImages []image.Image, err error) {
	if progress != nil {
		defer close(progress)
	}

	imageConverter := convert.NewImageConverter()

	asciiImages = make([]image.Image, 0, len(frames))
	for _, frame := range frames {
		charPixelMatrix := imageConverter.Image2CharPixelMatrix(frame, &convertOptions)
		asciiImage, err := drawer.DrawCharPixelMatrix2Image(charPixelMatrix, drawOptions)
		if err != nil {
			return nil, err
		}
		asciiImages = append(asciiImages, asciiImage)
		if progress != nil {
			progress <- 1
		}
	}
	return asciiImages, nil
}

// DrawCharPixelMatrix2Image draw a char pixel matrix to a image
func (drawer *ImageDrawer) DrawCharPixelMatrix2Image(charPixelMatrix [][]ascii.CharPixel, options DrawOptions) (img image.Image, err error) {
	drawFont, err := truetype.Parse(options.TTF)
	if err != nil {
		return nil, err
	}

	face := truetype.NewFace(drawFont, &truetype.Options{Size: options.FontSize})
	// get the font pixel width and height
	reflectFace := reflect.ValueOf(face)
	fontWidth := reflect.Indirect(reflectFace).FieldByName("maxw").Int()
	fontHeight := reflect.Indirect(reflectFace).FieldByName("maxh").Int()

	// check the matrix size
	matrixHeight := len(charPixelMatrix)
	if matrixHeight == 0 {
		return nil, errors.New("char matrix should not be empty")
	}
	matrixWidth := len(charPixelMatrix[0])
	dc := gg.NewContext(matrixWidth*int(fontWidth), matrixHeight*int(fontHeight))
	dc.SetFontFace(face)
	// set the background color
	backGroundColor := options.BackGroundColor
	dc.SetRGB(float64(backGroundColor.R)/255, float64(backGroundColor.G)/255, float64(backGroundColor.B)/255)
	dc.Clear()
	// set the font end color
	foreGroundColor := options.ForeGroundColor
	dc.SetRGB(float64(foreGroundColor.R)/255, float64(foreGroundColor.G/255), float64(foreGroundColor.B)/255)

	// draw the image
	for i, charPixelLine := range charPixelMatrix {
		for j, charPixel := range charPixelLine {
			if options.Colored {
				R, G, B, A := charPixel.R, charPixel.G, charPixel.B, charPixel.A
				dc.SetRGBA(float64(R)/255, float64(G)/255, float64(B)/255, float64(A)/255)
			}
			char := string(charPixel.Char)
			dc.DrawStringAnchored(char, float64(0+j*int(fontWidth)), float64(0+i*int(fontHeight)), 0, 1)
		}
	}
	return dc.Image(), nil
}

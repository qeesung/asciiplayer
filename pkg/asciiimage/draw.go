// asciiimage package convert the char pixel matrix that converted from
// the image2ascii to a png image that draw all ascii chars to the image
package asciiimage

import (
	"errors"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/qeesung/image2ascii/ascii"
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

type Drawer interface {
	DrawCharPixelMatrix2Image(charPixelMatrix [][]ascii.CharPixel, options DrawOptions) (img image.Image, err error)
}

type ImageDrawer struct {
}

func NewImageDrawer() Drawer {
	return &ImageDrawer{}
}

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
			fmt.Printf("T%+v, %+v\n", i, j)
			if options.Colored {
				R, G, B , A:= charPixel.R, charPixel.G, charPixel.B, charPixel.A
				dc.SetRGBA(float64(R)/255, float64(G)/255, float64(B)/255, float64(A)/255)
			}
			char := string(charPixel.Char)
			dc.DrawStringAnchored(char, float64(0+j*int(fontWidth)), float64(0+i*int(fontHeight)), 0, 1)
		}
	}
	return dc.Image(), nil
}

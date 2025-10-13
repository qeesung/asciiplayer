// Package render convert the char pixel matrix that converted from
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

// BatchConvertThenDraw batch draw the images and output the images
func (drawer *ImageDrawer) BatchConvertThenDraw(frames []image.Image,
	convertOptions convert.Options, drawOptions DrawOptions, progress chan<- int) (asciiImages []image.Image, err error) {
	if progress != nil {
		defer close(progress)
	}

	imageConverter := convert.NewImageConverter()

	asciiImages = make([]image.Image, 0, len(frames))
	for _, frame := range frames {
		charPixelMatrix := imageConverter.Image2CharPixelMatrix(frame, &convertOptions)
		asciiImage, _ := drawer.DrawCharPixelMatrix2Image(charPixelMatrix, drawOptions)
		asciiImages = append(asciiImages, asciiImage)
		if progress != nil {
			progress <- 1
		}
	}
	return asciiImages, nil
}

// 初始化绘图上下文及相关配置
func (drawer *ImageDrawer) initDrawContext(options DrawOptions) (dc *gg.Context, face truetype.Face, fontWidth, fontHeight int, err error) {
    // 解析字体
    drawFont, err := truetype.Parse(options.TTF)
    if err != nil {
        return nil, nil, 0, 0, err
    }

    // 创建字体Face
    face = truetype.NewFace(drawFont, &truetype.Options{Size: options.FontSize})
    
    // 通过反射获取字体尺寸
    reflectFace := reflect.ValueOf(face)
    fontWidth = int(reflect.Indirect(reflectFace).FieldByName("maxw").Int())
    fontHeight = int(reflect.Indirect(reflectFace).FieldByName("maxh").Int())

    return dc, face, fontWidth, fontHeight, nil
}

// 执行字符矩阵到图像的绘制
func (drawer *ImageDrawer) DrawCharPixelMatrix2Image(charPixelMatrix [][]ascii.CharPixel, options DrawOptions) (img image.Image, err error) {
    // 初始化绘图上下文
    dc, face, fontWidth, fontHeight, err := drawer.initDrawContext(options)
    if err != nil {
        return nil, err
    }

    // 检查矩阵有效性
    matrixHeight := len(charPixelMatrix)
    if matrixHeight == 0 {
        return nil, errors.New("char matrix should not be empty")
    }
    matrixWidth := len(charPixelMatrix[0])

    // 创建绘图上下文
    dc = gg.NewContext(matrixWidth*fontWidth, matrixHeight*fontHeight)
    dc.SetFontFace(face)

    // 设置背景色
    backGroundColor := options.BackGroundColor
    dc.SetRGB(
        float64(backGroundColor.R)/255,
        float64(backGroundColor.G)/255,
        float64(backGroundColor.B)/255,
    )
    dc.Clear()

    // 设置前景色
    foreGroundColor := options.ForeGroundColor
    dc.SetRGB(
        float64(foreGroundColor.R)/255,
        float64(foreGroundColor.G)/255, // 修复原代码中G通道的括号错误
        float64(foreGroundColor.B)/255,
    )

    // 绘制字符矩阵
    for i, charPixelLine := range charPixelMatrix {
        for j, charPixel := range charPixelLine {
            if !options.Colored {
                // 使用字符自带颜色
                dc.SetRGBA(
                    float64(charPixel.R)/255,
                    float64(charPixel.G)/255,
                    float64(charPixel.B)/255,
                    float64(charPixel.A)/255,
                )
            }
            dc.SetFontFace(face)
            char := string(charPixel.Char)
            dc.DrawStringAnchored(
                char,
                float64(j*fontWidth),
                float64(i*fontHeight),
                0, 1,
            )
        }
    }

    return dc.Image(), nil
}
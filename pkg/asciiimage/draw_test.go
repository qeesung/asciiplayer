package asciiimage

import (
	"github.com/qeesung/image2ascii/convert"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImageDrawer_DrawCharPixelMatrix2Image(t *testing.T) {
	assertions := assert.New(t)
	imageConverter := convert.NewImageConverter()
	imageFilename := "testdata/suolong.jpg"
	convertOptions := convert.DefaultOptions
	convertOptions.FitScreen = false
	convertOptions.Ratio = 0.1
	convertOptions.Colored = true
	charPixelMatrix := imageConverter.ImageFile2CharPixelMatrix(imageFilename, &convertOptions)

	drawer := NewImageDrawer()
	drawOptions := DefaultDrawOptions
	//drawOptions.Colored = false
	img, err := drawer.DrawCharPixelMatrix2Image(charPixelMatrix, drawOptions)
	assertions.Nil(err)
	assertions.NotNil(img)
}

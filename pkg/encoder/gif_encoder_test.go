package encoder

import (
	"fmt"
	"github.com/qeesung/image2ascii/convert"
	"github.com/stretchr/testify/assert"
	"image"
	"log"
	"os"
	"testing"
)

// TestGifEncoder_EncodeToFile test encode frames to gif
func TestGifEncoder_EncodeToFile(t *testing.T) {
	assertions := assert.New(t)
	frames := GetTheTestImages()
	// merge the frames
	gifEncoder := NewGifEncoder()
	err := gifEncoder.EncodeToFile("test.gif", frames, nil)
	assertions.Nil(err)
	// check the file
	_, err = os.Stat("test.gif")
	assertions.Nil(err)
	// remove the temp file
	os.Remove("test.gif")
}

// GetTheTestImages is util function to get the frames
func GetTheTestImages() []image.Image {
	imageFilenameList := make([]string, 0, 3)
	for i := 0; i < 3; i++ {
		imageFilenameList = append(imageFilenameList, fmt.Sprintf("testdata/suolong-%06d.png", i))
	}

	imageList := make([]image.Image, 0, 3)
	for i := 0; i < 3; i++ {
		imageFile, err := convert.OpenImageFile(imageFilenameList[i])
		if err != nil {
			log.Fatal(err)
		}
		imageList = append(imageList, imageFile)
	}
	return imageList
}

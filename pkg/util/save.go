// Package util contain some reused tool functions
package util

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

// SaveImageToFile is util function save the images to files
func SaveImageToFile(images []image.Image, filenamePrefix string) error {
	for index, img := range images {
		tempImageFilename := fmt.Sprintf("%s-%06d.png", filenamePrefix, index)
		outputFile, err := os.Create(tempImageFilename)
		if err != nil {
			return err
		}
		png.Encode(outputFile, img)
		outputFile.Close()
	}
	return nil
}

package util

import (
	"path/filepath"
)

// IsGif check if input file is gif file
func IsGif(filename string) bool {
	extension := filepath.Ext(filename)
	return extension == ".gif"
}

// IsSupportedImage check if the image type is supported type
func IsSupportedImage(filename string) bool {
	extension := filepath.Ext(filename)
	switch extension {
	case ".jpeg":
		fallthrough
	case ".jpg":
		fallthrough
	case ".png":
		return true
	default:
		return false
	}
}

// IsPng check if the input file is png file type
func IsPng(filename string) bool {
	extension := filepath.Ext(filename)
	return extension == ".png"
}

// IsJPG check if the input file is jpeg file type
func IsJPG(filename string) bool {
	extension := filepath.Ext(filename)
	return extension == ".jpg" || extension == ".jpeg"
}

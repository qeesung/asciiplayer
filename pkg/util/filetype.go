package util

import (
	"path/filepath"
)

func IsGif(filename string) bool {
	extension := filepath.Ext(filename)
	return extension == ".gif"
}

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

func IsPng(filename string) bool {
	extension := filepath.Ext(filename)
	return extension == ".png"
}

func IsJPG(filename string) bool {
	extension := filepath.Ext(filename)
	return extension == ".jpg" || extension == ".jpeg"
}

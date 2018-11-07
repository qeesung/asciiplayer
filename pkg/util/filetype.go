package util

import (
	"path/filepath"
)

func IsGif(filename string) bool {
	extension := filepath.Ext(filename)
	return extension == ".gif"
}

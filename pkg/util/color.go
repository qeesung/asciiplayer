package util

import (
	"gopkg.in/go-playground/colors.v1"
	"image/color"
)

// ConvertHexToRGB convert a color hex string to RGB color object
func ConvertHexToRGB(hexStr string) (color.RGBA, error) {
	hex, err := colors.ParseHEX(hexStr)
	if err == nil {
		return color.RGBA{
			R: hex.ToRGB().R,
			G: hex.ToRGB().G,
			B: hex.ToRGB().B,
			A: 255,
		}, nil
	}
	return color.RGBA{}, err
}

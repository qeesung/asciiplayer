package util

import (
	"gopkg.in/go-playground/colors.v1"
	"image/color"
)

func ConvertHexToRGB(hex string) (color.RGBA, error) {
	if hex, err := colors.ParseHEX(hex); err == nil {
		return color.RGBA{
			R: hex.ToRGB().R,
			G: hex.ToRGB().G,
			B: hex.ToRGB().B,
			A: 255,
		}, nil
	} else {
		return color.RGBA{}, err
	}
}

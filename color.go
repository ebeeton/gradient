// Package gradient defines functionality for generating linear color gradients.

package gradient

import (
	"encoding/hex"
	"errors"
	"fmt"
	"image/color"
	"math"
	"strings"
)

const (
	rgbByteLen  = 3
	rgbaByteLen = 4
)

// ColorFromHex converts a hex color string of the form #RRGGBB or #RRGGBBAA to
// a color.RGBA. If the former is provided alpha defaults to 255.
func ColorFromHex(h string) (color.RGBA, error) {
	data, err := hex.DecodeString(strings.Trim(h, "#"))
	if err != nil {
		return color.RGBA{}, err
	} else if len(data) == rgbByteLen {
		return color.RGBA{
			R: data[0],
			G: data[1],
			B: data[2],
			A: math.MaxUint8,
		}, nil
	} else if len(data) == rgbaByteLen {
		return color.RGBA{
			R: data[0],
			G: data[1],
			B: data[2],
			A: data[3],
		}, nil
	} else {
		return color.RGBA{}, errors.New("argument must be six or eight hex " +
			"characters")
	}
}

// ColorToHex converts a color.RGBA to a hex color string of the form #RRGGBB or
// #RRGGBBAA depending on the alpha parameter.
func ColorToHex(c color.RGBA, alpha bool) string {
	if alpha {
		return fmt.Sprintf("#%x%x%x%x", c.R, c.G, c.B, c.A)
	}

	return fmt.Sprintf("#%x%x%x", c.R, c.G, c.B)
}

func linearInterpolate(first, second color.RGBA, pos float64) color.RGBA {
	return color.RGBA{
		R: lerp(first.R, second.R, pos),
		G: lerp(first.G, second.G, pos),
		B: lerp(first.B, second.B, pos),
		A: lerp(first.A, second.A, pos),
	}
}

func lerp(first, second uint8, stop float64) uint8 {
	return uint8(float64(first)*(1.0-stop) + float64(second)*stop)
}

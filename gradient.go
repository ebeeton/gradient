// Package gradient defines functionality for generating linear color gradients.

package gradient

import (
	"image/color"
)

type stop struct {
	col color.RGBA
	pos float64
}

type gradientTable []stop

// GetGradient generates a gradient of count number of colors from a slice of
// Stops. The resulting slice of color.RGBA can be used as a palette.
func GetGradient(stops []Stop, count int) []color.RGBA {
	t := gradientTable{}
	p := []color.RGBA{}
	for _, s := range stops {
		c, err := colorFromHex(s.Color)
		if err != nil {
			panic("GetGradient: " + err.Error())
		}
		t = append(t, stop{col: c, pos: s.Position})
	}

	// Ensure that the values at either end of the gradient are the first and
	// last colors.
	step := 1.0 / float64(count-1)
	for i := 0; i < count; i++ {
		col := t.getInterpolatedColor(float64(i) * step)
		p = append(p, col)
	}
	return p
}

func (gt gradientTable) getInterpolatedColor(t float64) color.RGBA {
	for i := 0; i < len(gt)-1; i++ {
		c1 := gt[i]
		c2 := gt[i+1]
		if c1.pos <= t && t <= c2.pos {
			// Blend the two points we're between.
			t := (t - c1.pos) / (c2.pos - c1.pos)
			return linearInterpolate(c1.col, c2.col, t)
		}
	}

	// We're not between any points, so return the last color.
	return gt[len(gt)-1].col
}

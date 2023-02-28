// Package gradient defines functionality for generating linear color gradients.

package gradient

import "image/color"

// Stop is a hexadecimal color value and its position in a linear gradient from
// 0 to 1.
type Stop struct {
	Color    color.RGBA
	Position float64
}

# gradient

[![Go](https://github.com/ebeeton/gradient/actions/workflows/go.yml/badge.svg)](https://github.com/ebeeton/gradient/actions/workflows/go.yml)
[![CodeQL](https://github.com/ebeeton/gradient/actions/workflows/codeql.yml/badge.svg)](https://github.com/ebeeton/gradient/actions/workflows/codeql.yml)

A gradient that can generate slices of `color.RGBA` of arbitrary size.

The sample program below parses two colors from hex strings and uses them to
generate a gradient of nine colors.

```go
package main

import (
	"fmt"

	"github.com/ebeeton/gradient"
)

func main() {
	// Create two colors from hex strings of the form #RRGGBBAA.
	opaqueGreen, err := gradient.ColorFromHex("#00FF00FF")
	if err != nil {
		panic(err)
	}
	transparentMagenta, err := gradient.ColorFromHex("#FF00FF00")
	if err != nil {
		panic(err)
	}

	// Create a slice of stops with the colors at either "end".
	stops := []gradient.Stop{
		{Color: opaqueGreen, Position: 0.0},
		{Color: transparentMagenta, Position: 1.0},
	}

	// Get a gradient of nine colors and print them in hex. The second parameter
	// for ColorToHex includes the alpha channel value in the output.
	g := gradient.GetGradient(stops, 9)
	for _, c := range g {
		fmt.Println(gradient.ColorToHex(c, true))
	}
}
```

Output:

```
#00ff00ff
#1fdf1fdf
#3fbf3fbf
#5f9f5f9f
#7f7f7f7f
#9f5f9f5f
#bf3fbf3f
#df1fdf1f
#ff00ff00
```

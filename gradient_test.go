// Package gradient defines functionality for generating linear color gradients.

package gradient

import (
	"image/color"
	"testing"
)

func TestGetGradient(t *testing.T) {
	opaqueGreen, err := ColorFromHex("#00FF00FF")
	if err != nil {
		panic(err)
	}
	transparentMagenta, err := ColorFromHex("#FF00FF00")
	if err != nil {
		panic(err)
	}

	stops := []Stop{
		{Color: opaqueGreen, Position: 0.0},
		{Color: transparentMagenta, Position: 1.0},
	}
	wantColors := []string{
		"#00ff00ff",
		"#1fdf1fdf",
		"#3fbf3fbf",
		"#5f9f5f9f",
		"#7f7f7f7f",
		"#9f5f9f5f",
		"#bf3fbf3f",
		"#df1fdf1f",
		"#ff00ff00",
	}
	wantStops := 9
	g := GetGradient(stops, wantStops)

	if len(g) != wantStops {
		t.Errorf("Got stops %d, want %d.", len(g), wantStops)
	}

	for i, c := range g {
		var gotHex = ColorToHex(c, true)
		if gotHex != wantColors[i] {
			t.Errorf("Got %s, want %s.", gotHex, wantColors[i])
		}
	}
}

func TestGetInterpolatedColor(t *testing.T) {
	table := gradientTable{
		Stop{Color: color.RGBA{R: 0, G: 0, B: 0, A: 255}, Position: 0},
		Stop{Color: color.RGBA{R: 255, G: 255, B: 255, A: 255}, Position: 1},
	}
	want := color.RGBA{R: 127, G: 127, B: 127, A: 255}

	got := table.getInterpolatedColor(0.5)

	if got != want {
		t.Errorf("Got %v, want %v.", got, want)
	}
}

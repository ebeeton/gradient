// Package gradient defines functionality for generating linear color gradients.

package gradient

import (
	"image/color"
	"math"
	"testing"
)

func TestGetGradient(t *testing.T) {
	stops := []Stop{
		{Color: "#000000", Position: 0.0},
		{Color: "#FFFFFF", Position: 1.0},
	}

	wantStops := 3
	g := GetGradient(stops, wantStops)

	black := color.RGBA{R: 0, G: 0, B: 0, A: math.MaxUint8}
	grey := color.RGBA{R: math.MaxUint8 >> 1, G: math.MaxUint8 >> 1, B: math.MaxUint8 >> 1, A: math.MaxUint8}
	white := color.RGBA{R: math.MaxUint8, G: math.MaxUint8, B: math.MaxUint8, A: math.MaxUint8}

	if len(g) != wantStops {
		t.Errorf("Got stops %d, want %d.", len(g), wantStops)
	} else if g[0] != black {
		t.Errorf("Got black %v, want %v.", g[0], black)
	} else if g[1] != grey {
		t.Errorf("Got grey %v, want %v.", g[1], grey)
	} else if g[2] != white {
		t.Errorf("Got white %v, want %v.", g[2], white)
	}
}

func TestGetInterpolatedColor(t *testing.T) {
	table := gradientTable{
		stop{col: color.RGBA{R: 0, G: 0, B: 0, A: 255}, pos: 0},
		stop{col: color.RGBA{R: 255, G: 255, B: 255, A: 255}, pos: 1},
	}
	want := color.RGBA{R: 127, G: 127, B: 127, A: 255}

	got := table.getInterpolatedColor(0.5)

	if got != want {
		t.Errorf("Got %v, want %v.", got, want)
	}
}

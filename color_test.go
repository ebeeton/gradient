// Package gradient defines functionality for generating linear color gradients.

package gradient

import (
	"fmt"
	"image/color"
	"testing"
)

func TestColorFromHexRGB(t *testing.T) {
	hexColor := "#eb4034"
	want := color.RGBA{R: 235, G: 64, B: 52, A: 255}

	got, err := ColorFromHex(hexColor)
	if err != nil {
		t.Error(err.Error())
	} else if want != got {
		t.Errorf("Got %v, want %v.", got, want)
	}
}

func TestColorFromHexRGBA(t *testing.T) {
	hexColor := "#eb403480"
	want := color.RGBA{R: 235, G: 64, B: 52, A: 128}

	got, err := ColorFromHex(hexColor)
	if err != nil {
		t.Error(err.Error())
	} else if want != got {
		t.Errorf("Got %v, want %v.", got, want)
	}
}

func TestColorToHexRGB(t *testing.T) {
	color := color.RGBA{R: 4, G: 255, B: 67, A: 43}
	want := "#04ff43"

	got := ColorToHex(color, false)
	if want != got {
		t.Errorf("Got %s, want %s.", got, want)
	}
}

func TestColorToHexRGBA(t *testing.T) {
	color := color.RGBA{R: 17, G: 37, B: 167, A: 77}
	want := "#1125a74d"

	got := ColorToHex(color, true)
	if want != got {
		t.Errorf("Got %s, want %s.", got, want)
	}
}

func TestColorFromHexError(t *testing.T) {
	hexColor := "#eb403"

	_, err := ColorFromHex(hexColor)
	if err == nil {
		t.Error("Want error, got nil.")
	}
}

func TestLerp(t *testing.T) {
	var tests = []struct {
		first, second uint8
		stop          float64
		want          uint8
	}{
		{64, 192, 0.5, 128},
		{0, 255, 0.25, 63},
		{0, 255, 0.75, 191},
		{0, 255, 0, 0},
		{0, 255, 1.0, 255},
		{255, 255, 0.5, 255},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d,%f", tt.first, tt.second, tt.stop)
		t.Run(testname, func(t *testing.T) {
			got := lerp(tt.first, tt.second, tt.stop)
			if got != tt.want {
				t.Errorf("Got %d, want %d.", got, tt.want)
			}
		})
	}
}

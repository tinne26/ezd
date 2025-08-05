package ezd

import (
	"image"
	"testing"
)

func TestAnchorTranslateInt(t *testing.T) {
	rect10 := image.Rect(0, 0, 10, 10)
	rect9 := image.Rect(0, 0, 9, 9)
	rect10_9 := image.Rect(0, 0, 10, 9)

	tests := []struct {
		From       Anchor
		To         Anchor
		Rect       image.Rectangle
		OutX, OutY int
	}{
		{From: TopLeft, To: TopRight, Rect: rect10, OutX: 10, OutY: 0},
		{From: TopLeft, To: TopRight, Rect: rect9, OutX: 9, OutY: 0},
		{From: TopLeft, To: Center, Rect: rect10, OutX: 5, OutY: 5},
		{From: BottomRight, To: Center, Rect: rect10, OutX: -5, OutY: -5},

		{From: TopLeft, To: Center, Rect: rect9, OutX: 4, OutY: 4},
		{From: BottomRight, To: Center, Rect: rect9, OutX: -5, OutY: -5},

		{From: Center, To: CenterLeft, Rect: rect9, OutX: -4, OutY: 0},
		{From: Center, To: Center, Rect: rect9, OutX: 0, OutY: 0},
		{From: Center, To: CenterRight, Rect: rect9, OutX: 5, OutY: 0},

		{From: Center, To: TopCenter, Rect: rect9, OutX: 0, OutY: -4},
		{From: Center, To: Center, Rect: rect9, OutX: 0, OutY: 0},
		{From: Center, To: BottomCenter, Rect: rect9, OutX: 0, OutY: 5},

		{From: BottomLeft, To: TopRight, Rect: rect9, OutX: 9, OutY: -9},
		{From: Center, To: TopRight, Rect: rect9, OutX: 5, OutY: -4},
		{From: Center, To: TopCenter, Rect: rect9, OutX: 0, OutY: -4},
		{From: Center, To: TopRight, Rect: rect10_9, OutX: 5, OutY: -4},
		{From: Center, To: BottomCenter, Rect: rect9, OutX: 0, OutY: 5},
	}

	for i, test := range tests {
		x, y := test.From.TranslateInt(test.Rect, test.To)
		if x != test.OutX || y != test.OutY {
			t.Fatalf("test#%d: (%v) expected (x, y) = (%d, %d), got (%d, %d)", i, test, test.OutX, test.OutY, x, y)
		}
	}
}

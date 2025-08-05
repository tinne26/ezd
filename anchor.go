package ezd

import (
	"image"
	"strconv"
)

// Anchor references the point of an image to be
// aligned at some specific (x, y) coordinate pair.
type Anchor int

const (
	TopLeft      Anchor = 0b00_00
	TopCenter    Anchor = 0b00_01
	TopRight     Anchor = 0b00_10
	CenterLeft   Anchor = 0b01_00
	Center       Anchor = 0b01_01
	CenterRight  Anchor = 0b01_10
	BottomLeft   Anchor = 0b10_00
	BottomCenter Anchor = 0b10_01
	BottomRight  Anchor = 0b10_10
)

func (a Anchor) String() string {
	switch a {
	case TopLeft:
		return "TopLeft"
	case TopCenter:
		return "TopCenter"
	case TopRight:
		return "TopRight"
	case CenterLeft:
		return "CenterLeft"
	case Center:
		return "Center"
	case CenterRight:
		return "CenterRight"
	case BottomLeft:
		return "BottomLeft"
	case BottomCenter:
		return "BottomCenter"
	case BottomRight:
		return "BottomRight"
	default:
		return "Anchor(" + strconv.Itoa(int(a)) + ")"
	}
}

// TranslateF64 returns the translation offset from the current anchor to
// newAnchor for the given bounds.
//
// Low-level method for advanced use cases, rarely needed by ezd users.
func (a Anchor) TranslateF64(bounds image.Rectangle, newAnchor Anchor) (float64, float64) {
	// between a single axis of two aligns, there are only 3 possible
	// displacements: no move, half size (+/-), full size (+/-)
	var displacement = [5]float64{-1.0, -0.5, 0.0, 0.5, 1.0}
	horzDispl := float64(bounds.Dx()) * displacement[2+((newAnchor&0b11)-(a&0b11))]
	vertDispl := float64(bounds.Dy()) * displacement[2+((newAnchor>>2)-(a>>2))]
	return horzDispl, vertDispl
}

// TranslateF32 returns the translation offset from the current anchor to
// newAnchor for the given bounds.
//
// Low-level method for advanced use cases, rarely needed by ezd users.
func (a Anchor) TranslateF32(bounds image.Rectangle, newAnchor Anchor) (float32, float32) {
	var displacement = [5]float32{-1.0, -0.5, 0.0, 0.5, 1.0}
	horzDispl := float32(bounds.Dx()) * displacement[2+((newAnchor&0b11)-(a&0b11))]
	vertDispl := float32(bounds.Dy()) * displacement[2+((newAnchor>>2)-(a>>2))]
	return horzDispl, vertDispl
}

// TranslateInt returns the approximate translation offset from the current anchor
// to newAnchor for the given bounds. Anchors involving centered positions can have
// lossy translations, so [TranslateF64]() is recommended for translation chains.
// Alternatively, this is safe if bounds are even-sized.
//
// Low-level method for advanced use cases, rarely needed by ezd users.
func (a Anchor) TranslateInt(bounds image.Rectangle, newAnchor Anchor) (int, int) {
	w, h := bounds.Dx(), bounds.Dy()
	horzMove := (newAnchor & 0b11) - (a & 0b11)
	vertMove := (newAnchor >> 2) - (a >> 2)
	if a == Center && a != newAnchor {
		w += (w & 1) * int(horzMove)
		h += (h & 1) * int(vertMove)
	}

	var displacement = [5]int{-2, -1, 0, 1, 2}
	horzDispl := (w * displacement[2+horzMove]) >> 1
	vertDispl := (h * displacement[2+vertMove]) >> 1
	return horzDispl, vertDispl
}

package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tinne26/ezd"
)

func main() {
	image := ebiten.NewImage(96, 96)
	image.Fill(color.White)
	if err := ebiten.RunGame(&App{image: image, x: 32, y: 32}); err != nil {
		log.Fatal(err)
	}
}

type App struct {
	drawer ezd.Drawer
	anchor ezd.Anchor // use arrow keys to change
	image  *ebiten.Image
	rads   float64 // use mouse wheel to change
	x, y   int
}

func (a *App) Layout(w, h int) (int, int) {
	return w, h
}

func (a *App) Update() error {
	preAnchor := a.anchor
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyEscape):
		return ebiten.Termination
	case ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft):
		a.x, a.y = ebiten.CursorPosition()
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft):
		a.anchor = (a.anchor & 0b11_00) | max((a.anchor&0b00_11)-1, 0)
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowRight):
		a.anchor = (a.anchor & 0b11_00) | min((a.anchor&0b00_11)+1, 2)
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowUp):
		a.anchor = (max(((a.anchor&0b11_00)>>2)-1, 0) << 2) | (a.anchor & 0b00_11)
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowDown):
		a.anchor = (min(((a.anchor&0b11_00)>>2)+1, 2) << 2) | (a.anchor & 0b00_11)
	}
	if a.anchor != preAnchor {
		a.rads = 0.0
		xOffset, yOffset := preAnchor.TranslateInt(a.image.Bounds(), a.anchor)
		a.x += xOffset
		a.y += yOffset
	}

	_, yScroll := ebiten.Wheel()
	scrollFactor := 0.1
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		scrollFactor = 0.025
	}
	a.rads = math.Mod(a.rads+yScroll*scrollFactor, 2*math.Pi)

	ebiten.SetWindowTitle(fmt.Sprintf(
		"drawing image's %s at (%d, %d) (rotated %.02f degrees)",
		a.anchor.String(), a.x, a.y, ezd.Rads2Degs(a.rads),
	))

	return nil
}

func (a *App) Draw(canvas *ebiten.Image) {
	canvas.Fill(color.Black)

	// draw image using current anchor
	a.drawer.Filter = ebiten.FilterLinear // optional
	a.drawer.DrawRotatedAt(a.image, a.anchor, canvas, float64(a.x), float64(a.y), a.rads)

	// draw anchor point as a 4x4 filled area
	anchorRect := image.Rect(a.x-2, a.y-2, a.x+2, a.y+2)
	canvas.SubImage(anchorRect).(*ebiten.Image).Fill(color.RGBA{255, 0, 255, 255})
}

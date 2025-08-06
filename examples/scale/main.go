package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tinne26/ezd"
)

func main() {
	image := ebiten.NewImage(96, 96)
	image.Fill(color.White)
	if err := ebiten.RunGame(&App{image: image, x: 32, y: 32, scale: 1.0}); err != nil {
		log.Fatal(err)
	}
}

type App struct {
	drawer ezd.Drawer
	anchor ezd.Anchor // use arrow keys to change
	image  *ebiten.Image
	scale  float64
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
		// NOTE: this adjustment is not exact, but it's ok for the example
		xOffset, yOffset := preAnchor.Translate(a.image.Bounds(), a.anchor)
		a.x += int(xOffset * a.scale)
		a.y += int(yOffset * a.scale)
	}

	_, yScroll := ebiten.Wheel()
	scrollFactor := 0.025
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		scrollFactor = 0.005
	}
	a.scale = max(a.scale+yScroll*scrollFactor, 0.1)

	titleInfo := fmt.Sprintf("drawing image's %s at (%d, %d) (scale x%.02f)", a.anchor.String(), a.x, a.y, a.scale)
	ebiten.SetWindowTitle(titleInfo)

	return nil
}

func (a *App) Draw(canvas *ebiten.Image) {
	canvas.Fill(color.Black)

	// draw image using current anchor
	a.drawer.Filter = ebiten.FilterLinear
	a.drawer.DrawScaledAt(a.image, a.anchor, canvas, float64(a.x), float64(a.y), a.scale)

	// draw anchor point as a 4x4 filled area
	anchorRect := image.Rect(a.x-2, a.y-2, a.x+2, a.y+2)
	canvas.SubImage(anchorRect).(*ebiten.Image).Fill(color.RGBA{255, 0, 255, 255})
}

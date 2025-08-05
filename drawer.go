package ezd

import "github.com/hajimehoshi/ebiten/v2"

type Drawer struct {
	ebiten.DrawImageOptions
}

// DrawAtF64 draws source's anchor to target's (x, y). For example, if anchor is [BottomRight],
// DrawAtF64 will draw source such that its bottom right corner is at (x, y).
//
// When working with float64 positions, you might want to set Drawer.Filter = ebiten.FilterLinear.
func (d *Drawer) DrawAtF64(source *ebiten.Image, anchor Anchor, target *ebiten.Image, x, y float64) {
	trx, try := anchor.TranslateF64(source.Bounds(), TopLeft)
	memo := d.GeoM
	d.GeoM.Translate(x+trx, y+try)
	target.DrawImage(source, &d.DrawImageOptions)
	d.GeoM = memo
}

// DrawAt draws source's anchor to target's (x, y). For example, if anchor is [Center],
// DrawAt will draw source such that its center is at (x, y).
//
// Notice that when working with images with odd sizes and center anchors, the positions
// might be truncated. If you need more precision, consider [Drawer.DrawAtF64]().
func (d *Drawer) DrawAt(source *ebiten.Image, anchor Anchor, target *ebiten.Image, x, y int) {
	trx, try := anchor.TranslateInt(source.Bounds(), TopLeft)
	memo := d.GeoM
	d.GeoM.Translate(float64(x+trx), float64(y+try))
	target.DrawImage(source, &d.DrawImageOptions)
	d.GeoM = memo
}

// DrawRotatedAtF64 draws source's anchor to target's (x, y) applying the given
// rotation.
//
// When working with rotations, you might want to set Drawer.Filter = ebiten.FilterLinear.
func (d *Drawer) DrawRotatedAtF64(source *ebiten.Image, anchor Anchor, target *ebiten.Image, x, y, rads float64) {
	rx, ry := anchor.TranslateF64(source.Bounds(), TopLeft)
	memo := d.GeoM
	d.GeoM.Translate(rx, ry)
	d.GeoM.Rotate(rads)
	d.GeoM.Translate(x, y)
	target.DrawImage(source, &d.DrawImageOptions)
	d.GeoM = memo
}

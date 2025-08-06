package ezd

import "github.com/hajimehoshi/ebiten/v2"

type Drawer struct {
	ebiten.DrawImageOptions
}

// DrawAtInt draws source's anchor to target's (x, y). For example, if anchor is [Center],
// DrawAtInt will draw source such that its center is at (x, y).
//
// Notice that when working with images with odd sizes and center anchors, the positions
// might be truncated. If you need more precision, consider [Drawer.DrawAt]().
func (d *Drawer) DrawAtInt(source *ebiten.Image, anchor Anchor, target *ebiten.Image, x, y int) {
	trx, try := anchor.TranslateInt(source.Bounds(), TopLeft)
	memo := d.GeoM
	d.GeoM.Translate(float64(x+trx), float64(y+try))
	target.DrawImage(source, &d.DrawImageOptions)
	d.GeoM = memo
}

// DrawAt draws source's anchor to target's (x, y). For example, if anchor is [BottomRight],
// DrawAt will draw source such that its bottom right corner is at (x, y).
//
// When working with float64 positions, you might want to set Drawer.Filter = [ebiten.FilterLinear].
func (d *Drawer) DrawAt(source *ebiten.Image, anchor Anchor, target *ebiten.Image, x, y float64) {
	trx, try := anchor.Translate(source.Bounds(), TopLeft)
	memo := d.GeoM
	d.GeoM.Translate(x+trx, y+try)
	target.DrawImage(source, &d.DrawImageOptions)
	d.GeoM = memo
}

// DrawRotatedAt draws source's anchor to target's (x, y) applying the given
// rotation.
//
// When working with rotations, you might want to set Drawer.Filter = [ebiten.FilterLinear].
func (d *Drawer) DrawRotatedAt(source *ebiten.Image, anchor Anchor, target *ebiten.Image, x, y, rads float64) {
	trx, try := anchor.Translate(source.Bounds(), TopLeft)
	memo := d.GeoM
	d.GeoM.Translate(trx, try)
	d.GeoM.Rotate(rads)
	d.GeoM.Translate(x, y)
	target.DrawImage(source, &d.DrawImageOptions)
	d.GeoM = memo
}

// DrawScaledAt draws source's anchor to target's (x, y) with the given scale.
//
// When scaling images, you might want to set Drawer.Filter = [ebiten.FilterLinear].
func (d *Drawer) DrawScaledAt(source *ebiten.Image, anchor Anchor, target *ebiten.Image, x, y, scale float64) {
	trx, try := anchor.Translate(source.Bounds(), TopLeft)
	memo := d.GeoM
	d.GeoM.Translate(trx, try)
	d.GeoM.Scale(scale, scale)
	d.GeoM.Translate(x, y)
	target.DrawImage(source, &d.DrawImageOptions)
	d.GeoM = memo
}

// DrawGeoAt draws source's anchor to target's (x, y) with the given scale and rotation.
//
// When applying scaling and rotation, you might want to set Drawer.Filter = [ebiten.FilterLinear].
func (d *Drawer) DrawGeoAt(source *ebiten.Image, anchor Anchor, target *ebiten.Image, x, y, scale, rads float64) {
	trx, try := anchor.Translate(source.Bounds(), TopLeft)
	memo := d.GeoM
	d.GeoM.Translate(trx, try)
	d.GeoM.Scale(scale, scale)
	d.GeoM.Rotate(rads)
	d.GeoM.Translate(x, y)
	target.DrawImage(source, &d.DrawImageOptions)
	d.GeoM = memo
}

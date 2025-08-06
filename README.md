# ezd

[![Go Reference](https://pkg.go.dev/badge/github.com/tinne26/ezd.svg)](https://pkg.go.dev/github.com/tinne26/ezd)

Easy drawing on Ebitengine with anchors:

```Golang
// Drawer is just a wrapper over ebiten.DrawImageOptions
var drw ezd.Drawer

// You can access DrawImageOptions fields and methods directly
drw.Filter = ebiten.FilterLinear
drw.ColorScale.ScaleAlpha(0.8)

// But you also have comfy methods to draw.
// This method would be read like "draw source's
// bottom-center at target's (x, y)"
drw.DrawAt(source, ezd.BottomCenter, target, x, y)
```

This can help make your code much less verbose if you are drawing simple images at different positions based on anchors, including basic scaling and rotations.

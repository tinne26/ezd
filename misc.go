package ezd

import "math"

func Degs2Rads(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func Rads2Degs(rads float64) float64 {
	return rads * 180.0 / math.Pi
}

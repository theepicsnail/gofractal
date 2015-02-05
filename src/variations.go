package main

import "math"

type Variation func(x, y float64) (float64, float64)

func getR2(x, y float64) float64 {
	return x*x + y*y
}

var spherical = func(x, y float64) (float64, float64) {
	r := getR2(x, y)
	if r == 0 {
		return x, y
	}

	return x / r, y / r
}

var swirl = func(x, y float64) (float64, float64) {
	r := getR2(x, y)
	return x*math.Sin(r) - y*math.Cos(r),
		x*math.Cos(r) + y*math.Sin(r)
}

var scale = func(x, y float64) (float64, float64) {
	return x / 2, y / 2
}

var linear = func(x, y float64) (float64, float64) {
	return x, y
}
var sinusoidal = func(x, y float64) (float64, float64) {
	return math.Sin(x), math.Sin(y)
}

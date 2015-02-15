package main

import "github.com/theepicsnail/gofractal/flameutil"
import (
	"flag"
	"math"
)

var FLAG_PERCENT = flag.Float64("p", 0, "Percentage of animation [0,1]")

var (
	loop = func(phase float64) func(float64) float64 {
		return func(in float64) float64 {
			return (1 - math.Cos(phase+in*math.Pi*2)) / 2
		}
	}
)

func configure(config *flameutil.FlameConfig) {
	p := *FLAG_PERCENT

	config.AddVariation(.5, flameutil.Scale(loop(0)(p)))
	config.AddFlameFunction(.33, flameutil.AffineTransform(.5, 0, 0, 0, .5, 0))
	config.AddFlameFunction(.33, flameutil.AffineTransform(.5, 0, 1, 0, .5, 0))
	config.AddFlameFunction(.34, flameutil.AffineTransform(.5, 0, 0, 0, .5, 1))
}

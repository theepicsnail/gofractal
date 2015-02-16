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

	config.AddVariation(.99-loop(0)(p)*.2, flameutil.Spherical)
	config.AddVariation(loop(0)(p)*.2+.01, flameutil.Heart)

	config.AddFlameFunction(.2,
		//flameutil.Rotation(p*math.Pi*2),
		flameutil.AffineTransform(-1, 0, 0, 0, 1, 0),
		flameutil.Color_NOOP)
	config.AddFlameFunction(.4,
		flameutil.AffineTransform(2, 0, 0, 0, 2, 0),
		flameutil.Color_BLEND(flameutil.Color_RED))
	config.AddFlameFunction(.4,
		flameutil.AffineTransform(.5, 0, 0, 0, .5, 0),
		flameutil.Color_BLEND(flameutil.Color_WHITE))

}

package main

import "math"
import "./variations"

var (
	linear = func(start, end float64) func(float64) float64 {
		return func(in float64) float64 {
			return start + (end-start)*in
		}
	}
	loop = func(phase float64, freq int) func(float64) float64 {
		return func(in float64) float64 {
			return math.Sin(phase+float64(freq)*in*math.Pi*2)/2 + .5
		}
	}
)

func configure(p float64) {

	//Variations = append(Variations, WeightedVariation{loop(0, 1)(p), swirl})
	///Variations = append(Variations, WeightedVariation{loop(math.Pi/3*2, 1)(p), sinusoidal})
	//Variations = append(Variations, WeightedVariation{loop(math.Pi/3*4, 1)(p), spherical})

	Variations = append(Variations, WeightedVariation{0.5, variations.Sinusoidal})
	Variations = append(Variations, WeightedVariation{0.5, variations.Swirl})

	angle := p * math.Pi * 2
	sin := math.Sin(angle)
	cos := math.Cos(angle)
	Coefs = append(Coefs, Coef{.2, cos, -sin, 0, sin, cos, 0})
	Coefs = append(Coefs, Coef{.4, 2, 0, 0, 0, 2, 0})
	Coefs = append(Coefs, Coef{.4, .5, 0, 0, 0, .5, 0})

}

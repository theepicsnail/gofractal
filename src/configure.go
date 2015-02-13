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
	/*
		var moveTowards = func(tx, ty float64) func(float64, float64) (float64, float64) {
			return func(x, y float64) (float64, float64) {
				return (tx + x) / 2, (ty + y) / 2
			}
		}
			Variations = append(Variations, WeightedVariation{.33, moveTowards(0, 1)})
			Variations = append(Variations, WeightedVariation{.34, moveTowards(-1, -1)})
			Variations = append(Variations, WeightedVariation{.34, moveTowards(1, -1)})
			Coefs = append(Coefs, Coef{1, 1, 0, 0, 0, 1, 0})
			return*/

	//Variations = append(Variations, WeightedVariation{loop(0, 1)(p), swirl})
	//Variations = append(Variations, WeightedVariation{loop(math.Pi/3*2, 1)(p), sinusoidal})
	//Variations = append(Variations, WeightedVariation{loop(math.Pi/3*4, 1)(p), spherical})
	/*
		if p < .5 {
			Variations = append(Variations, WeightedVariation{linear(0, 1)(p * 2), variations.Sinusoidal})
			Variations = append(Variations, WeightedVariation{linear(1, 0)(p * 2), variations.Spherical})
		} else {
			p = (p - .5) * 2
			Variations = append(Variations, WeightedVariation{linear(1, 0)(p), variations.Sinusoidal})
			Variations = append(Variations, WeightedVariation{linear(0, 1)(p), variations.Swirl})
		}
	*/
	Variations = append(Variations, WeightedVariation{1, variations.Scale(p * 2)})
	/*angle := math.Pi / 3
	sin := math.Sin(angle)
	cos := math.Cos(angle)
	Coefs = append(Coefs, Coef{.2, cos, -sin, 0, sin, cos, 0})

	Coefs = append(Coefs, Coef{.5, 2, 0, 0, 0, 2, 0})
	Coefs = append(Coefs, Coef{.5, .5, 0, 0, 0, .5, 0})
	*/

	Coefs = append(Coefs, Coef{.5, .5, 0, 0, 0, .5, 0})
	Coefs = append(Coefs, Coef{.25, .5, 0, 1, 0, .5, 0})
	Coefs = append(Coefs, Coef{.25, .5, 0, 0, 0, .5, 1})

}

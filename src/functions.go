package main

import (
	"math"
	"math/rand"
)

type Coef struct{ p, a, b, c, d, e, f float64 }

type WeightedVariation struct {
	w float64
	f Variation
}

var rnd = rand.New(rand.NewSource(99))

var Coefs = [...]Coef{
	//    p  a  b  c  d  e  f
	// Rotation matrix
	Coef{.3, math.Cos(0), -math.Sin(0), 0, math.Sin(0), math.Cos(0), 0},
	Coef{.4, 2, 0, 0, 0, 2, 0},  // J = 0
	Coef{.4, -1, 0, 0, 0, 1, 0}, // J = 1
	//Coef{.7, -1, 0, 1, 0, -1, 1}, // J = 1
}
var Variations = []WeightedVariation{
	{.4, swirl},
	{.2, sinusoidal},
	{0, linear},
	{.4, spherical},
}

func configure(p float64) {
	/*
		linear := func(start, end float64) func(float64) float64 {
			return func(in float64) float64 {
				return start + (end-start)*in
			}
		}
	*/
	loop := func(phase float64, freq int) func(float64) float64 {
		return func(in float64) float64 {
			return math.Sin(phase+float64(freq)*in*math.Pi*2)/2 + .5
		}
	}

	/*
		Variations[0].w = linear(0, .5)(p)
		Variations[1].w = linear(1, .2)(p)
		Variations[2].w = linear(0, .0)(p)
		Variations[3].w = linear(0, .3)(p)
	*/

	Variations[0].w = (loop(0, 1)(p)) / 2
	Variations[1].w = (loop(2, 1)(p)) / 2
	//Variations[2].w = (1 - loop(0, 1)(p)) / 2
	Variations[3].w = (1 - loop(4, 1)(p)) / 2
}

// F_j(x,y) = Sum(V_k in Variations, ... )
func F(j int, x, y float64) (float64, float64) {
	var outX float64
	var outY float64

	// X and y coords passed into the variation
	var coef = Coefs[j]
	vx := coef.a*x + coef.b*y + coef.c
	vy := coef.d*x + coef.e*y + coef.f

	for _, variation := range Variations {
		resultX, resultY := variation.f(vx, vy)
		outX += variation.w * resultX
		outY += variation.w * resultY
	}

	return outX, outY
}

func randomFunction() func(float64, float64) (float64, float64) {
	r := rnd.Float64()
	j := 0
	for ; r < 1; j++ {
		r += Coefs[j].p
		//fmt.Println("j:", j, "r:", r)
	}
	j -= 1

	return func(x, y float64) (float64, float64) {
		return F(j, x, y)
	}
}

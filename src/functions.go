package main

import (
	"./variations"
	"math"
	"math/rand"
)

type Coef struct{ p, a, b, c, d, e, f float64 }

type WeightedVariation struct {
	w float64
	f variations.Variation
}

var rnd = rand.New(rand.NewSource(99))

var angle = math.Pi
var Coefs = []Coef{}
var Variations = []WeightedVariation{}

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

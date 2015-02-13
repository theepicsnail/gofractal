package main

import (
	"./variations"
	"math"
	"math/rand"
)

// Replace this with probability, point->point func, color -> color func
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
func F(j int, p Point) Point {
	out := Point{}

	// This should just be the application of an affine transform
	var coef = Coefs[j]
	vx := coef.a*p.x + coef.b*p.y + coef.c
	vy := coef.d*p.x + coef.e*p.y + coef.f

	for _, variation := range Variations {
		resultX, resultY := variation.f(vx, vy)
		out.x += variation.w * resultX
		out.y += variation.w * resultY
	}

	return out
}

func randomFunction() func(Point) Point {
	// Rand int % length(Coefs) ?
	r := rnd.Float64()
	j := 0
	for ; r < 1; j++ {
		r += Coefs[j].p
		//fmt.Println("j:", j, "r:", r)
	}
	j -= 1

	// Precompute and store these functions?
	// Perhaps they are inlined since it's just tail recursion.
	return func(p Point) Point {
		return F(j, p)
	}
}

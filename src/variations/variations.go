package variations

import "math"

type Variation func(x, y float64) (float64, float64)

func getR2(x, y float64) float64 {
	return x*x + y*y
}
func getR(x, y float64) float64 {
	return math.Sqrt(getR2(x, y))
}
func invsqrt(f float64) float64 {
	if f == 0 {
		return 0
	}
	return 1 / math.Sqrt(f)
}

var Scale = func(factor float64) Variation {
	return func(x, y float64) (float64, float64) {
		return x * factor, y * factor
	}
}

var Sinusoidal = func(x, y float64) (float64, float64) {
	return math.Sin(x), math.Sin(y)
}

var Spherical = func(x, y float64) (float64, float64) {
	r := getR2(x, y)
	if r == 0 {
		return x, y
	}

	return x / r, y / r
}

var Swirl = func(x, y float64) (float64, float64) {
	r := getR2(x, y)
	return x*math.Sin(r) - y*math.Cos(r),
		x*math.Cos(r) + y*math.Sin(r)
}

var Horseshoe = func(x, y float64) (float64, float64) {
	inv_r := invsqrt(getR2(x, y))
	return inv_r * (x - y) * (x + y), inv_r * 2 * x * y
}

// Polar

// Handkerchief

// Heart

// Disc

// Spiral

// Hyperbolic

// Diamond

// Ex

// Julia

var Bent = func(x, y float64) (float64, float64) {
	if x >= 0 {
		x *= 2
	}
	if y < 0 {
		y /= 2
	}
	return x, y
}

// Waves

var Fisheye = func(x, y float64) (float64, float64) {
	return Eyefish(y, x)
}

// Popcorn

var Exponential = func(x, y float64) (float64, float64) {
	var exp = math.Exp(x - 1)
	var ypi = math.Pi * y
	return exp * math.Cos(ypi), exp * math.Sin(ypi)
}

// Power

var Cosine = func(x, y float64) (float64, float64) {
	return math.Cos(math.Pi*x) * math.Cosh(y), -math.Sin(math.Pi*x) * math.Sinh(y)
}

// Rings

// Fan

// Blob

// PDJ

// Fan2

// Rings2

var Eyefish = func(x, y float64) (float64, float64) {
	var coef = 2 / (getR(x, y) + 1)
	return x * coef, y * coef
}

var Bubble = func(x, y float64) (float64, float64) {
	var coef = 4 / (getR2(x, y) + 4)
	return coef * x, coef * y
}

var Cylinder = func(x, y float64) (float64, float64) {
	return math.Sin(x), y
}

// Perspectice

// Noise

// JuliaN

// JuliaScope

// Blur

// Gaussian

// RadialBlur

// Pie

// Ngon

// Curl

// Rectangles

// Arch

var Tangent = func(x, y float64) (float64, float64) {
	return math.Sin(x) / math.Cos(y), math.Tan(y)
}

// Square

// Rays

// Blade

// Secant

// Twintrian

var Cross = func(x, y float64) (float64, float64) {
	var coef = 1 / math.Abs(x*x-y*y)
	return coef * x, coef * y
}

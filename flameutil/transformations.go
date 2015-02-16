package flameutil

import "math"

type Transformation func(*Point)
type Variation struct {
	weight         float64
	transformation Transformation
}
type FlameFunction struct {
	probability    float64
	transformation Transformation
	colorMutator   ColorTransformation
}

var (
	// Affine transformations
	Rotation = func(angle float64) Transformation {
		c := math.Cos(angle)
		s := math.Sin(angle)
		return func(p *Point) {
			p.X, p.Y = p.X*c-p.Y*s, p.X*s+p.Y*c
		}
	}

	Scale = func(factor float64) Transformation {
		return func(p *Point) {
			p.X *= factor
			p.Y *= factor
		}
	}

	AffineTransform = func(a, b, c, d, e, f float64) Transformation {
		return func(p *Point) {
			p.X, p.Y = a*p.X+b*p.Y+c, d*p.X+e*p.Y+f
		}
	}

	// Independant flame transformations
	Sinusoidal = func(p *Point) {
		p.X = math.Sin(p.X)
		p.Y = math.Sin(p.Y)
	}

	Spherical = func(p *Point) {
		if r := p.R2(); r != 0 {
			p.X /= r
			p.Y /= r
		}
	}

	Swirl = func(p *Point) {
		angle := p.R2()
		Rotation(angle)(p)
	}

	Heart = func(p *Point) {
		angle := p.Angle() * p.R()
		p.X = math.Sin(angle)
		p.Y = -math.Cos(angle)
	}
)

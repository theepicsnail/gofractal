package main

import (
	"image"
	"image/color"
	//"image/draw"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

const (
	WIDTH            = 1200
	HEIGHT           = 800
	XMIN     float64 = -2
	XMAX     float64 = 1
	YMIN     float64 = -1
	YMAX     float64 = 1
	MAX_ITER float64 = 100 // Actually an int. but go sucks
)

//https://code.google.com/p/gorilla/source/browse/color/hsv.go?r=ef489f63418265a7249b1d53bdc358b09a4a2ea0
func HSVToRGB(h, s, v float64) (r, g, b uint8) {
	var fR, fG, fB float64
	i := math.Floor(h * 6)
	f := h*6 - i
	p := v * (1.0 - s)
	q := v * (1.0 - f*s)
	t := v * (1.0 - (1.0-f)*s)
	switch int(i) % 6 {
	case 0:
		fR, fG, fB = v, t, p
	case 1:
		fR, fG, fB = q, v, p
	case 2:
		fR, fG, fB = p, v, t
	case 3:
		fR, fG, fB = p, q, v
	case 4:
		fR, fG, fB = t, p, v
	case 5:
		fR, fG, fB = v, p, q
	}
	r = uint8((fR * 255) + 0.5)
	g = uint8((fG * 255) + 0.5)
	b = uint8((fB * 255) + 0.5)
	return
}

func calculateColor(col *color.RGBA, x, y float64) {
	z := complex(x, y)
	c := complex(x, y)

	var iter float64 = 0
	for ; (iter < MAX_ITER) && (cmplx.Abs(z) < 2); iter++ {
		z = z*z + c
	}

	if iter == MAX_ITER {
		col.R, col.G, col.B = 0, 0, 0
	} else {
		col.R, col.G, col.B = HSVToRGB(iter/MAX_ITER, 1, 1)
	}
	//real(z), 1, imag(z))

}

func main() {
	// Create an image.
	m := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))

	c := color.RGBA{255, 255, 255, 255}
	/*
		//Fill it with black
		draw.Draw(m, m.Bounds(),
			&image.Uniform{color.RGBA{0, 0, 0, 255}},
			image.ZP, draw.Src)
	*/

	// draw the pixels!
	for x := float64(m.Bounds().Min.X); x < float64(m.Bounds().Max.X); x++ {
		for y := float64(m.Bounds().Min.Y); y < float64(m.Bounds().Max.Y); y++ {
			coord_x := XMIN + (XMAX-XMIN)*x/WIDTH
			coord_y := YMIN + (YMAX-YMIN)*y/HEIGHT
			calculateColor(&c, coord_x, coord_y)
			m.Set(int(x), int(y), c)
		}
	}

	w, _ := os.Create("new.png")
	defer w.Close()
	png.Encode(w, m) //Encode writes the Image m to w in PNG format.

}

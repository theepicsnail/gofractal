package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"math"
	"path/filepath"
	//"image/draw"
	"image/png"
	//"math"
	//"math/cmplx"
	"os"
)

const (
	WIDTH  = 400
	HEIGHT = 400
)

func PixelToPoint(x, y int) (float64, float64) {
	return float64(x*2)/float64(WIDTH) - 1,
		float64(y*2)/float64(HEIGHT) - 1
}

func PointToPixel(x, y float64) (int, int) {
	return int((x + 1) / 2 * WIDTH),
		int((y + 1) / 2 * HEIGHT)
}

func main() {

	var percent *float64
	var dir *string
	percent = flag.Float64("p", 0, "")
	dir = flag.String("dir", ".", "")
	flag.Parse()
	configure(*percent)

	//Coefs[0].p = *percent
	//Coefs[1].p = 1 - *percent

	//Coefs[0].p = .1
	//Coefs[1].p = .9

	//Variations[0].w = *percent
	//Variations[1].w = 1 - *percent

	histogram := make([][]float64, HEIGHT)
	for i := range histogram {
		histogram[i] = make([]float64, WIDTH)
	}

	var max_val float64 = 0
	/*for py := 0; py < HEIGHT; py++ {
		for px := 0; px < WIDTH; px++ {
			//p, ok := swirl[spherical[Point{x, y}]]

			x, y := PointToPixel(
				spherical(swirl(
					spherical(swirl(
						PixelToPoint(px, py))))))

			if x < 0 || y < 0 || x >= WIDTH || y >= HEIGHT {
				continue
			}

			histogram[y][x] += 1
			if histogram[y][x] > max_val {
				max_val = histogram[y][x]
			}
		}
	}*/

	for iter := 0; iter < 100; iter++ {
		x, y := rnd.Float64(), rnd.Float64() // TODO(snail) pick a random point for this.

		for i := 0; i < 10000; i++ {
			x, y = randomFunction()(x, y)
			if i > 20 {
				// Plot
				px, py := PointToPixel(x, y)
				if px < 0 || py < 0 || px >= WIDTH || py >= HEIGHT {
					continue
				}

				// Keep a running max
				histogram[py][px] += 1
				if histogram[py][px] > max_val {
					max_val = histogram[py][px]
				}
			}
		}
	}
	// Create an image.
	m := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))

	c := color.RGBA{255, 255, 255, 255}
	// draw the pixels!
	//for x := float64(m.Bounds().Min.X); x < float64(m.Bounds().Max.X); x++ {
	//	for y := float64(m.Bounds().Min.Y); y < float64(m.Bounds().Max.Y); y++ {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			c.R = uint8(math.Log(histogram[y][x]) * 255 / math.Log(max_val))
			c.G = c.R
			c.B = c.R
			m.Set(int(x), int(y), c)
		}
	}

	path := filepath.Join(*dir, fmt.Sprintf("img%.2f.png", *percent))
	fmt.Println(path)
	w, _ := os.Create(path)
	//	fmt.Sprintf("%v/img%.2f.png", filepath.Abs(*dir), *percent))
	defer w.Close()
	png.Encode(w, m) //Encode writes the Image m to w in PNG format.

}

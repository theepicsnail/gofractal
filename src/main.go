package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"path/filepath"
)

const (
	WIDTH  = 800
	HEIGHT = 800

	SCALE = 2 // How large of a coordinate system to fit in the image 2 = (-2 to 2)
	XMIN  = -1 * SCALE
	XMAX  = 1 * SCALE
	YMIN  = -1 * SCALE
	YMAX  = 1 * SCALE

	X_DELTA = XMAX - XMIN
	Y_DELTA = YMAX - YMIN
)

/* Utilities for point <-> pixel based on the above constants */
func PixelToPoint(x, y int) (float64, float64) {
	return float64(x)*X_DELTA/float64(WIDTH) + XMIN,
		float64(y)*Y_DELTA/float64(HEIGHT) + YMIN
}

func PointToPixel(x, y float64) (int, int) {
	return int((x - XMIN) / X_DELTA * WIDTH),
		int((y - YMIN) / Y_DELTA * HEIGHT)
}

func main() {

	// setup
	var percent *float64
	var dir *string
	percent = flag.Float64("p", 0, "")
	dir = flag.String("dir", ".", "")
	flag.Parse()
	configure(*percent)

	histogram := make([][]float64, HEIGHT)
	for i := range histogram {
		histogram[i] = make([]float64, WIDTH)
	}

	var max_val float64 = 0

	// is 100 different starting points any better than just 1 starting point?
	for iter := 0; iter < 1; iter++ {
		x, y := rnd.Float64(), rnd.Float64()
		for i := 0; i < 10000000; i++ { // Points per random walk
			x, y = randomFunction()(x, y)

			if i > 20 {
				// Give the random walk time to get within a pixel of accurancy.

				px, py := PointToPixel(x, y)
				// TODO this check could be done before converting using px,py vs {X,Y}{MIN,MAX}
				if px < 0 || py < 0 || px >= WIDTH || py >= HEIGHT {
					continue
				}

				// Keep a running max
				// Perhaps this could be quicker instead of 3*2 lookups
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
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			// fix brightness
			c.R = uint8(math.Log(histogram[y][x]) * 255 / math.Log(max_val))
			c.G = c.R/2 + c.R/4
			c.B = 0 //c.R
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

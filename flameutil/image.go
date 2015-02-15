package flameutil

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

var FLAG_IMAGE_WIDTH = flag.Int("width", 100, "Width of the output image.")
var FLAG_IMAGE_HEIGHT = flag.Int("height", 100, "Height of the output image.")

// Super sampling
// Desnsity estimation

/*
	This struct holds the information frequired to render an image. This means
	it's more than width*height*color. It also stores a histogram of locations,
	the internal width/height may be different from the end image width and
	height due to super sampling, etc...

	There is the method Save(filepath) which will actually ddraw the image with
	the specified width/height/options.
*/
type FlameImage struct {
	imageWidth      int
	imageHeight     int
	histogramWidth  int
	histogramHeight int
	histogram       [][]histogramData
	histogramMax    float64
}

type histogramData struct {
	count float64
	// color
	// log desnitry thing for bluring?
}

func createHistogram(w, h int) [][]histogramData {
	histogram := make([][]histogramData, h) // 2d slices of data
	data := make([]histogramData, h*w)      // underlying data structure
	for row := range histogram {
		histogram[row], data = data[:w], data[w:]
		// It's okay that we're losing the original slice reference to
		// the entire underlying data structure, we don't need it as
		// we really only need the histogram model of the data.
	}
	return histogram
}

func NewImage() *FlameImage {
	/* For now image{Width,Height} == histogram{Width,Height}
	When I add in super sampling this will nolonger be true */

	w := *FLAG_IMAGE_WIDTH
	h := *FLAG_IMAGE_HEIGHT

	return &FlameImage{
		// Image size
		w, h,
		// Histogram size
		w, h,
		createHistogram(w, h),
		0,
	}
}

/* Adds a point to the histogram */
func (img *FlameImage) addPoint(point *Point) {
	px := int((point.X + 1) / 2 * float64(img.imageWidth))
	py := int((point.Y + 1) / 2 * float64(img.imageHeight))
	if px >= 0 && py >= 0 && px < img.imageWidth && py < img.imageWidth {
		img.histogram[py][px].count++
		possibleMax := img.histogram[py][px].count
		if possibleMax > img.histogramMax {
			img.histogramMax = possibleMax
		}
	}
}

func (flame *FlameImage) Save(file string) {
	img := image.NewRGBA(image.Rect(0, 0, flame.imageWidth, flame.imageHeight))
	c := color.RGBA{255, 255, 255, 255}

	// fix brightness

	for y := 0; y < flame.imageHeight; y++ {
		for x := 0; x < flame.imageWidth; x++ {
			brightness := math.Log(flame.histogram[y][x].count) /
				math.Log(flame.histogramMax)
			c.R = uint8(brightness * 255)
			c.G = c.R/2 + c.R/4
			c.B = 0 //c.R
			img.Set(int(x), int(y), c)
		}
	}

	w, _ := os.Create(file)
	defer w.Close()
	png.Encode(w, img)
}

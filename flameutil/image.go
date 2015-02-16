package flameutil

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

var FLAG_IMAGE_WIDTH = flag.Int("width", 400, "Width of the output image.")
var FLAG_IMAGE_HEIGHT = flag.Int("height", 400, "Height of the output image.")
var FLAG_SUPER_SAMPLING = flag.Bool("supersample", true, "Enable supersampling (3x)")
var FLAG_DENSITY_ESTIMATION = flag.Bool("density_estimation", true, "Enable density estimation")

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
	samplingsize    int
}

type histogramData struct {
	count float64
	color color.RGBA
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
	sampling := 1
	if *FLAG_SUPER_SAMPLING {
		sampling = 3
	}
	return &FlameImage{
		// Image size
		w, h,
		// Histogram size
		w * sampling, h * sampling,
		createHistogram(w*sampling, h*sampling),
		0,
		sampling,
	}
}

/* Adds a point to the histogram */
func (img *FlameImage) addPoint(point *Point, color color.RGBA) {
	px := int((point.X + 1) / 2 * float64(img.histogramWidth))
	py := int((point.Y + 1) / 2 * float64(img.histogramHeight))
	if px >= 0 && py >= 0 && px < img.histogramWidth && py < img.histogramWidth {
		img.histogram[py][px].count++
		img.histogram[py][px].color = color
		possibleMax := img.histogram[py][px].count
		if *FLAG_SUPER_SAMPLING {
			possibleMax = 0
			baseX := px / img.samplingsize * img.samplingsize
			baseY := py / img.samplingsize * img.samplingsize
			for x := baseX; x < baseX+img.samplingsize; x++ {
				for y := baseY; y < baseY+img.samplingsize; y++ {
					possibleMax += img.histogram[y][x].count
				}
			}
		}
		if possibleMax > img.histogramMax {
			img.histogramMax = possibleMax
		}
	}
}

func copy_color(color color.RGBA) color.RGBA {
	return color
}

func (flame *FlameImage) Save(file string) {
	img := image.NewRGBA(image.Rect(0, 0, flame.imageWidth, flame.imageHeight))

	for y := 0; y < flame.imageHeight; y++ {
		for x := 0; x < flame.imageWidth; x++ {
			img.Set(x, flame.imageHeight-y-1, flame.compute_color(x, y))
		}
	}

	w, _ := os.Create(file)
	defer w.Close()
	png.Encode(w, img)
}

func (flame *FlameImage) compute_color(x, y int) color.RGBA {
	pixel_color := color.RGBA{0, 0, 0, 255}
	count := 0.0

	if *FLAG_SUPER_SAMPLING {
		sampling := flame.samplingsize
		// average the supersampling cells
		tot_r := 0.0
		tot_g := 0.0
		tot_b := 0.0
		tot_samples := 0.0
		for dx := 0; dx < sampling; dx++ {
			for dy := 0; dy < sampling; dy++ {
				c := flame.histogram[y*sampling+dy][x*sampling+dx].color
				count := flame.histogram[y*sampling+dy][x*sampling+dx].count
				tot_samples += count
				tot_r += float64(c.R) * count
				tot_g += float64(c.G) * count
				tot_b += float64(c.B) * count
			}
		}
		if tot_samples == 0 {
			pixel_color = color.RGBA{0, 0, 0, 255}
		} else {
			pixel_color = color.RGBA{
				uint8(tot_r / tot_samples),
				uint8(tot_g / tot_samples),
				uint8(tot_b / tot_samples),
				255,
			}

			if 199 < x && x < 202 && 199 < y && y < 205 {
				fmt.Println(pixel_color)
				/*				pixel_color.R = 255
								pixel_color.G = 255
								pixel_color.B = 255
								tot_samples = flame.histogramMax*/
			}
		}
		count = tot_samples
	} else {
		pixel_color = copy_color(flame.histogram[y][x].color)
		count = flame.histogram[y][x].count
	}

	// Fix brightness
	brightness := math.Log(count+1) / math.Log(flame.histogramMax+1)
	pixel_color.R = uint8(brightness * float64(pixel_color.R))
	pixel_color.G = uint8(brightness * float64(pixel_color.G))
	pixel_color.B = uint8(brightness * float64(pixel_color.B))
	pixel_color.A = uint8(255)

	if !*FLAG_DENSITY_ESTIMATION {
		return pixel_color
	}
	/*
		I HAVE NO IDEA HOW TO DO DENSITY ESTIMATION

		max_blur_radius := 3
		blur_radius := int(float64(max_blur_radius) / math.Pow(flame.histogram[y][x].count+1.0, 1.0))
		//- int(float64(max_blur_radius)*brightness) //flame.histogram[y][x].count/flame.histogramMax)
		total_points := 0.0
		total_r := 0.0
		total_g := 0.0
		total_b := 0.0

		for nx := x - blur_radius; nx <= x+blur_radius; nx++ {
			if nx < 0 || nx >= flame.imageWidth {
				continue
			}
			for ny := y - blur_radius; ny <= y+blur_radius; ny++ {
				if ny < 0 || ny >= flame.imageHeight {
					continue
				}
				c := flame.histogram[ny][nx].color
				cnt := flame.histogram[ny][nx].count
				total_points += 1
				if cnt != 0 {
					d_sq := 1.0 //.75 * (1.0 - (float64((y-ny)*(y-ny)+(x-nx)*(x-nx)))/float64(blur_radius*blur_radius))
					total_r += float64(c.R) * d_sq
					total_g += float64(c.G) * d_sq
					total_b += float64(c.B) * d_sq
				}
			}
		}
		c.R = uint8(total_r / total_points * brightness)
		c.G = uint8(total_g / total_points * brightness)
		c.B = uint8(total_b / total_points * brightness)
		if blur_radius == max_blur_radius {
			//		c.G = uint8(255)
		}
		return c*/
	return pixel_color
}

package main

import "flag"
import "github.com/theepicsnail/gofractal/flameutil"

func main() {
	flag.Parse()
	config := flameutil.NewFlameConfig()
	configure(config)

	image := flameutil.NewImage()
	flameutil.Render(config, image)
	image.Save("out.png")
}

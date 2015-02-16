package main

import (
	"flag"
	"fmt"
)
import "github.com/theepicsnail/gofractal/flameutil"

var FLAG_FILE = flag.String("file", "img%.2f.png", "Filename to save image as, use %.2f for percenage")

func main() {
	flag.Parse()
	config := flameutil.NewFlameConfig()
	configure(config)

	image := flameutil.NewImage()
	flameutil.Render(config, image)
	image.Save(fmt.Sprintf(*FLAG_FILE, *FLAG_PERCENT))
	fmt.Println(*FLAG_PERCENT)
}

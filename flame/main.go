package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)
import "github.com/theepicsnail/gofractal/flameutil"

var FLAG_FILE = flag.String("file", "img%.2f.png", "Filename to save image as, use %.2f for percenage")

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	config := flameutil.NewFlameConfig()
	configure(config)

	image := flameutil.NewImage()
	flameutil.Render(config, image)
	file := fmt.Sprintf(*FLAG_FILE, *FLAG_PERCENT)
	image.Save(file)
	fmt.Println(file)
}

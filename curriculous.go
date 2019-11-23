package main

import (
	"curriculous/gif"
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

const font = "/usr/share/fonts/comic-sans/design.graffiti.comicsansms.ttf"

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

	var list = []string{"pipi", "caca", "prout", "cul", "nichon", "caca", "prout", "cul", "nichon", "caca", "prout", "cul", "nichon", "caca", "prout", "cul", "nichon", "caca", "prout", "cul", "nichon", "caca", "prout", "cul", "FIN"}

	var filename = "foo.gif"
	var font_color = [3]uint8{0, 0, 0}
	var back_color = [3]uint8{255, 255, 255}
	var opts = gif.BuildOptions(font, font_color, back_color, filename)
	gif.MakeGif(list, opts)
}

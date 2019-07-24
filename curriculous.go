package main

import (
	"curriculous/gif"
)

const font = "/usr/share/fonts/comic-sans/design.graffiti.comicsansms.ttf"


func main() {


	var list = []string{"pipi", "caca", "prout", "cul", "nichon"}

	var font_color = [3]uint8{0, 0, 0}
	var back_color = [3]uint8{255, 255, 255}
	var opts = gif.BuildOptions(font, font_color, back_color)
	gif.MakeGif(list, opts)
}

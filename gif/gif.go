package gif

import (
	"bufio"
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
)

func addText(dst *image.Paletted, text string, opts Options) {
	mask := textMask(text, opts.FontPath, dst.Bounds().Dx(), dst.Bounds().Dy())
	im := image.NewPaletted(dst.Bounds(), palette.Plan9)
	fillRect(im, 0, 0, 400, 300, opts.FontColor)
	draw.DrawMask(dst, dst.Bounds(), im, image.ZP, mask, image.ZP, draw.Over)
}

func textMask(text, FontPath string, width, height int) *image.Alpha {
	dc := gg.NewContext(width, height)

	w, h := float64(width/2), float64(height/2)
	// draw text
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace(FontPath, 128)
	dc.DrawStringAnchored(text, w, h, 0.5, 0.5)
	return dc.AsMask()
}

type Options struct {
	FontColor color.RGBA
	BackColor color.RGBA
	FontPath  string
}

func BuildOptions(font string, bcolor, fcolor [3]uint8) Options {
	return Options{
		FontPath:  font,
		FontColor: color.RGBA{fcolor[0], fcolor[1], fcolor[2], 255},
		BackColor: color.RGBA{bcolor[0], bcolor[1], bcolor[2], 255},
	}
}

func MakeGif(list []string, opts Options) error {
	var frames = []*image.Paletted{}

	// Creating delay array
	var delay = []int{}

	for _, word := range list {
		frames = appendFrame(frames, word, opts)
		delay = append(delay, 50)
	}

	res := gif.GIF{
		Image:     frames,
		Delay:     delay,
		LoopCount: 0,
		Disposal:  nil,
		Config: image.Config{
			// ColorModel: color.RGBAModel,
			Width:  400,
			Height: 300,
		},
		BackgroundIndex: 0,
	}

	f, err := os.Create("foo.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	err = gif.EncodeAll(w, &res)
	if err != nil {
		panic(err)
	}
	w.Flush()

	return nil
}

// func addFrame(dst *image.Paletted, word string, frames []*image.Paletted{}) []*image.Paletted{} {
func appendFrame(frames []*image.Paletted, word string, opts Options) []*image.Paletted {
	fmt.Println(word)
	dst := image.NewPaletted(image.Rect(0, 0, 400, 300), palette.Plan9)
	fillRect(dst, 0, 0, 400, 300, opts.BackColor)
	addText(dst, word, opts)

	return append(frames, dst)
}

func fillRect(dst *image.Paletted, x, y, w, h int, color color.RGBA) {
	for i := 0; i <= w; i++ {
		for j := 0; j <= h; j++ {
			dst.Set(x+i, y+j, color)
		}
	}
}

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,                        // фон — черный
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // линия — ярко-зеленая
}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	f, err := os.Create("out.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lissajous(f)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // количество колебаний
		res     = 0.001 // разрешение
		size    = 100   // размер изображения
		nframes = 64    // количество кадров
		delay   = 8     // задержка между кадрами
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // случайная частота
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

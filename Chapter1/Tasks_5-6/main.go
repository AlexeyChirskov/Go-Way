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

// var palette = []color.Color{color.White, color.Black}

var palette = []color.Color{color.RGBA{0, 0, 0, 1}, color.RGBA{34, 255, 0, 1}, color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
	/*Упражнение 1.5. Измените палитру программы lissajous так, чтобы изобра­жение было зеленого цвета на черном фоне,
	чтобы быть более похожим на экран ос­ циллографа. Чтобы создать веб-цвет #RRGGBB, воспользуйтесь
	инструкцией color.RGBA{0xRRj0xGG,0xBB,0xff}, в которой каждая пара шестнадцатеричных цифр представляет
	яркость красного, зеленого и синего компонентов пикселя.*/
	// var palette = []color.Color{color.RGBA{0, 0, 0, 1}, color.RGBA{34, 255, 0, 1}}

	/*Упражнение 1.6. Измените программу lissajous так, чтобы она генерировала изображения разных цветов,
	добавляя в палитру palette больше значений, а затем выводя их путем изменения третьего аргумента функции
	SetColorlndex некоторым нетривиальным способом.*/
	// uint8(rand.Intn(len(palette)))

}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(len(palette))))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

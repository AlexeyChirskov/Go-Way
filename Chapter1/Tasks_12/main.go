package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	count   int
	palette = []color.Color{color.RGBA{0, 0, 0, 1}, color.RGBA{34, 255, 0, 1}, color.White, color.Black}
	cycles  = 5
	res     = 0.001
	size    = 100
	nframes = 64
	delay   = 8
)

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	// http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var err error

		for k, v := range r.URL.Query() {
			fmt.Println(k, " => ", v)
			switch k {
			case "cycles":
				cycles, err = strconv.Atoi(v[0])
				if err != nil {
					log.Println(err)
					return
				}
			case "res":
				res, err = strconv.ParseFloat(v[0], 64)
				if err != nil {
					log.Println(err)
					return
				}
			case "size":
				size, err = strconv.Atoi(v[0])
				if err != nil {
					log.Println(err)
					return
				}
			case "nframes":
				nframes, err = strconv.Atoi(v[0])
				if err != nil {
					log.Println(err)
					return
				}
			case "delay":
				delay, err = strconv.Atoi(v[0])
				if err != nil {
					log.Println(err)
					return
				}
			}
		}

		lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)

	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer) {
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2.0*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(rand.Intn(len(palette))))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

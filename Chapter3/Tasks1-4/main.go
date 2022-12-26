package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // Размер канвы в пикселях
	cells         = 100                 // Количество ячеек сетки
	xyrange       = 30.0                // Диапазон осей // (-xyrange..+ xyrange)
	xyscale       = width / 2 / xyrange // Пикселей в единице х или у
	zscale        = height * 0.4        // Пикселей в единице z
	angle         = math.Pi / 6         // Углы осей х, у (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30°),cos(30°)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		surface(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surface(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width=,%d' height='%d' >", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, isValid1 := corner(i+1, j)
			bx, by, isValid2 := corner(i, j)
			cx, cy, isValid3 := corner(i, j+1)
			dx, dy, isValid4 := corner(i+1, j+1)
			if !isValid1 || !isValid2 || !isValid3 || !isValid4 {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)

		}
		fmt.Fprintf(out, "</svg>")
	}
}

func corner(i, j int) (float64, float64, bool) {
	// Ищем угловую точку (x,y) ячейки (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхности z
	z := f(x, y)
	// Изометрически проецируем (x,y,z) на двумерную канву SVG (sx,sy)
	sx := width/2 + (x+y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	if math.IsNaN(sx) || math.IsNaN(sy) {
		return sx, sy, false
	}
	return sx, sy, true
}

func f(x, у float64) float64 {
	r := math.Hypot(x, у) // Расстояние от (0,0)
	return math.Sin(r) / r
}

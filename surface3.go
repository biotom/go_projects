package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: grey; fill: white; stroke-width: 0.7' "+"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, colour, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, _, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, _, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, _, err := corner(i+1, j+1)
			if err != nil {
				continue
			}

			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' "+"style='fill:%s' />\n", ax, ay, bx, by, cx, cy, dx, dy, colour)

		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, string, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//z := f(x, y)
	z := f(x, y)
	colour := "#0000ff"

	if math.IsInf(z, 0) == true || math.IsNaN(z) == true {
		return 1, 1, "give it up", fmt.Errorf("nah")
	}

	if z >= 0 {
		colour = "#ff0000"
	} //else {
	//colour := "#0000ff"
	//}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, colour, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func heart(x, y float64) float64 {
	r := (math.Pow(math.Pow(x, 2)+math.Pow(y, 2)-1, 3) - math.Pow(x, 2)*math.Pow(y, 3))
	return r
}

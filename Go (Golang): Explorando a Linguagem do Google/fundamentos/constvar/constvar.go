package main

import "math"

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	//	Forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2)

	println("A área da circunferencia é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	println(a, b, c, d)

	var e, f bool = true, false
	println(e, f)

	g, h, i := 2, false, "epa!"
	println(g, h, i)
}

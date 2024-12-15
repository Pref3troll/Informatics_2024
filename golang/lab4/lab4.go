package lab4

import "math"

const a = 1.2
const b = 0.48

func calculateY(a, b, x float64) float64 {
	return ((math.Pow(b, 3) + math.Pow(math.Sin(a*x), 2)) / ((math.Acos(x * b * x)) + math.Pow(math.E, (-x/2))))
}

func TaskA(xn, xk, xdel float64) []float64 {
	var yValues []float64
	for x := xn; x <= xk; x += xdel {
		yValues = append(yValues, calculateY(a, b, x))
	}
	return yValues
}

func TaskB(xv []float64) []float64 {
	var yValues []float64
	for _, x := range xv {
		yValues = append(yValues, calculateY(a, b, x))
	}
	return yValues
}

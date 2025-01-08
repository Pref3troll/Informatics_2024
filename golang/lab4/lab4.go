package lab4

import (
	"fmt"
	"math"
)

func calculateY(a, b, x float64) float64 {
	return ((math.Pow(b, 3) + math.Pow(math.Sin(a*x), 2)) / ((math.Acos(x * b * x)) + math.Pow(math.E, (-x/2))))
}

func TaskA(a, b, xn, xk, xdel float64) []float64 {
	var yValues []float64
	for x := xn; x <= xk; x += xdel {
		yValues = append(yValues, calculateY(a, b, x))
	}
	return yValues
}

func TaskB(a, b float64, xv []float64) []float64 {
	var yValues []float64
	for _, x := range xv {
		yValues = append(yValues, calculateY(a, b, x))
	}
	return yValues
}

func Lab4() {
	fmt.Println("Задача А:\n", TaskA(1.2, 0.48, 0.7, 2.2, 0.3))
	fmt.Println("Задача В:\n", TaskB(1.2, 0.48, []float64{0.25, 0.36, 0.56, 0.94, 1.28}))
}

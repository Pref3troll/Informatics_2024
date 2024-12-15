package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

func main() {
	fmt.Println("Емельчиков Егор Павлович")
	fmt.Println("Задача А:\n", lab4.TaskA(0.7, 2.2, 0.3))
	fmt.Println("Задача B:\n", lab4.TaskB([]float64{0.25, 0.36, 0.56, 0.94, 1.28}))
}

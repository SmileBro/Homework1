package main

import (
	"fmt"
	"homework1/SqrEquation"
)

func main() {
	var a = new(float64)
	var b = new(float64)
	var c = new(float64)
	fmt.Println("Введите значения a,b,c:")
	fmt.Scan(a, b, c)

	x1, x2 := SqrEquation.Solve(*a, *b, *c)
	fmt.Printf("x1 = %f\nx2 = %f\n", x1, x2)
}

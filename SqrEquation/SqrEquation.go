package SqrEquation

import (
	"fmt"
	"homework1/Desc"
	"math"
)

func Solve(a, b, c float64) (float64, float64) {
	desc := Desc.Find(a, b, c)
	if desc < 0 {
		x1 := 0.0
		x2 := 0.0
		fmt.Println("Дискриминант < 0")
		return x1, x2
	} else {
		x1 := (-b - math.Sqrt(desc)) / (2.0 * a)
		x2 := (-b + math.Sqrt(desc)) / (2.0 * a)
		return x1, x2
	}
}

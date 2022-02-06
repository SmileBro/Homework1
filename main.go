package main

import (
	"fmt"
	"homework1/Bank"
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
	fmt.Println("Второе задание")
	fmt.Println("Введите сумму вклада")
	fmt.Scan(a)
	fmt.Println("Введите годовой процент вклада")
	fmt.Scan(b)
	fmt.Printf("Сумма вклада через 5 лет: %f\n", Bank.Deposit(*a, *b))
}

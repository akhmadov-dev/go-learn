package main

import "fmt"

// Oddiy arifmetik amallar
func main() {
	var a float64 = 10
	var b float64 = 20

	c := a + b
	d := a - b
	e := a * b
	y := a / b

	fmt.Println("a + b =", c)
	fmt.Println("a - b =", d)
	fmt.Println("a * b =", e)
	fmt.Println("a / b =", y)
}
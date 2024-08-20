package main

import "fmt"

func main() {
	var in1 float32
	var in2 float32

	fmt.Print("Birinchi sonni kiriting: ")
	fmt.Scanln(&in1)
	fmt.Print("Ikkinchi sonni kiriting: ")
	fmt.Scanln(&in2)

	result := summ(in1, in2)

	fmt.Printf("Natija: %v\n", result)
}

func summ(a float32, b float32) float32 {
	return a + b;
}
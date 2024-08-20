package main

import "fmt"

// foydalanuvchi soni bo'yicha sikl
func main() {
	var input int

	fmt.Print("Natural son kiriting: ")
	fmt.Scanln(&input)

	for i := 0; i <= input; i++ {
		fmt.Printf("- %v\n", i)
	}
}
package main

import "fmt"

// oddiy shart operatorlari
// sonning musbat, manfiy yoki nol ekanligini aniqlash
func main() {
	var input float64
	var result string

	fmt.Print("Son kiriting: ")
	fmt.Scanln(&input)

	if input > 0 {
		result = "musbat"
	} else if input < 0 {
		result = "manfiy"
	} else {
		result = "nol"
	}

	fmt.Printf("Ushbu son - %s\n", result)
}
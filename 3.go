package main

import "fmt"

// Kiritish va chiqarish
func main() {
	var input string

	// foydalanuvchidan ma'lumot olish
	fmt.Print("Iltimos, ismingizni kiriting: ")
	fmt.Scanln(&input)


	// kiritilgan ma'lumotni ekranga chiqarish
	fmt.Printf("Salom, %s!\n", input)
}
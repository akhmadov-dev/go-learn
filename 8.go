package main

import "fmt"

// Toq va juft sonlar
func main() {
	n := 20
	var checkNumber string

	for i:=1; i<=n; i++ {
		if i % 2 == 0 {
			checkNumber = "juft"
		} else {
			checkNumber = "toq"
		}

		fmt.Printf("%v - %s\n", i, checkNumber)
	}
}
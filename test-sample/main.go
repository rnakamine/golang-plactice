package main

import "fmt"

func checkEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var n int

	n = 10
	fmt.Println(checkEven(n))

	n = 11
	fmt.Println(checkEven(n))
}

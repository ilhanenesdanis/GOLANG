package main

import "fmt"

func main() {
	a := 10
	b := 10

	if a > b {
		fmt.Println("Büyüktür")
	} else if b == a {
		fmt.Println("Eşittir")
	} else {
		fmt.Println("Küçüktür")
	}
}

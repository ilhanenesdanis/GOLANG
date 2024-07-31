package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(12, 33))

	a, b := swap("Hello", "Golang")

	fmt.Println(a, b)
}

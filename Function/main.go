package main

import "fmt"

func sayHello() {
	println("Hello")
}
func sayHelloWithMessage(message string) {
	println(message)
}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func multipleAdd(terms ...int) (int, int) {
	result := 0

	for _, term := range terms {
		result += term
	}
	return len(terms), result
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

func main() {

	sayHello()

	sayHelloWithMessage("Hello go")

	fmt.Println(add(12, 33))

	a, b := swap("Hello", "Golang")

	fmt.Println(a, b)

	numTerms, sum := multipleAdd(1, 2, 3, 4, 5, 6, 6, 7, 8, 345, 123, 33)

	fmt.Println(numTerms, sum)

	//anonymous function

	addFunc := func(terms ...int) (numTerms int, sum int) {
		for _, term := range terms {
			sum += term
		}
		numTerms = len(terms)

		return
	}

	num, sm := addFunc(2, 4)
	println(num, sm)

	defer fmt.Println("Hello")

	fmt.Println("Go")

}

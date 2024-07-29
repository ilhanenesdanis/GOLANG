package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	fmt.Println(a)
	myArr1 := [3]int{}

	myArr1[0] = 32
	myArr1[1] = 22
	myArr1[2] = 41

	fmt.Println(myArr1)

	//colors

	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println(colors)

	var numbers = [5]int{1, 2, 3, 4, 52}

	fmt.Println(numbers)

	fmt.Println("Number of numbers ", len(numbers))

	myArr2 := [...]int{1, 3, 4, 5, 6, 2, 4}
	fmt.Println(myArr2)

	var cars [3]string
	cars[0] = "BMW"
	cars[1] = "Honda"
	cars[2] = "Audi"

	for i, value := range cars {
		fmt.Println("i= ", i, " &Value = ", value)
	}
}

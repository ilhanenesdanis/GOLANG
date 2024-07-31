package main

import "fmt"

func main() {

	primes := [6]int{2, 3, 4, 5, 6, 4}

	var s []int = primes[1:4]

	fmt.Println(s)

	names := [4]string{
		"İlhan",
		"Rümeysa",
		"Ramazan",
		"Hakan",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b)

	b[0] = "xxx"

	fmt.Println(a, b)
	fmt.Println(names)

}

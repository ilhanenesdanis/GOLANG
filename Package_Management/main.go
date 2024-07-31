package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	//fmt

	fmt.Println("Hello go")

	fmt.Println("Random value :", rand.Intn(10))

	fmt.Println(strings.Contains("Test", "es"))

	fmt.Println(strings.Count("test", "t"))

	fmt.Println(strings.HasPrefix("unit test", "unit"))

	fmt.Println(strings.HasPrefix("dosya.rar", "rar"))

	fmt.Println(strings.Index("test", "e"))

}

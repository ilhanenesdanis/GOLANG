package main

import (
	"fmt"
	"strconv"
)

func main() {
	//convert

	var myStyring = "1"
	var x = 10

	var f float32 = 2.0

	fmt.Println(myStyring, x, f)

	//string to int

	number, _ := strconv.Atoi(myStyring)

	fmt.Println(number)

	result := number + 2

	fmt.Println(result)
}

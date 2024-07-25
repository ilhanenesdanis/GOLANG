package main

import "fmt"

func main() {
	//konsol veri çıkışı

	str := "The quick red fox"

	str2 := "jumped over"

	str3 := "the lazy brown dog "

	aNumber := 58

	//	isTrue := true

	stringLength, _ := fmt.Println(str, str2, str3)

	fmt.Println("String length:", stringLength)

	fmt.Printf("Value of aNumber: %v\n", aNumber)

}

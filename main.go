package main

import "fmt"

const thy = "Türk Hava Yolları"
const pgs = "Pegasus"

//constant enumaration
type Brand string

const (
	FACEBOOK  Brand = "Facebook"
	MICROSOFT Brand = "Microsoft"
)

func main() {
	//Constants

	fmt.Println(thy)

	fmt.Println(FACEBOOK)

	fmt.Println(MICROSOFT)

}

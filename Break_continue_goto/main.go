package main

import "fmt"

func main() {

	//i := 0
	// for {
	// 	if i >= 3 {
	// 		break
	// 	}

	// 	fmt.Println("i'nin değeri : ", i)

	// 	i++

	// }

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Çıktı : ", i)
	}

}

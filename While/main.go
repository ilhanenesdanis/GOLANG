package main

import (
	"fmt"
	"time"
)

func main() {
	//for (while)

	sum := 1

	for sum < 10000 {
		sum += sum
		fmt.Println(sum)
		time.Sleep(300 * time.Millisecond)
	}

}

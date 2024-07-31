package main

import "fmt"

type MyClass struct {
	Lat, Long float64
}

var m map[string]MyClass

func main() {
	//1
	m = make(map[string]MyClass)

	m["MyLocation"] = MyClass{
		40.68433, -74.39967,
	}
	fmt.Println(m["MyLocation"])
	//2 key value pair
	mayMap := make(map[int]string)
	fmt.Println(mayMap)

	mayMap[43] = "Foo"
	mayMap[12] = "Bar"

	fmt.Println(mayMap)

	//3

	states := make(map[string]string)
	states["Ist"] = "İstanbul"
	states["Ank"] = "Ankara"
	states["Svs"] = "Sivas"

	fmt.Println(states)

	sivas := states["Svs"]

	fmt.Println("Seçili Şehir : ", sivas)

	delete(states, "Ank")

	fmt.Println(states)
}

package main

import "fmt"

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

//Constructor
//varsayılan ve boş constructor
func NewHuman() *Human {
	h := new(Human)

	return h
}

func NewHumanWithParam(firstName, lastName string) *Human {
	h := new(Human)
	h.FirstName = firstName
	h.LastName = lastName
	return h
}

func main() {

	x := Human{
		FirstName: "İlhan",
		LastName:  "Daniş",
		Age:       22,
	}
	fmt.Println(x.FirstName)
	fmt.Println(x.LastName)
	fmt.Println(x.Age)

	a := new(Human)

	a.FirstName = "Ramazan"
	fmt.Println(a.FirstName)

	//constructor

	y := NewHuman()

	y.Age = 51

	fmt.Println(y.Age)

	z := NewHumanWithParam("İlhan", "Daniş")
	fmt.Println(z.FirstName)
	fmt.Println(z.LastName)
}

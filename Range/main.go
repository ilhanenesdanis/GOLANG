package main

func main() {

	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		println("Array Item :", a[i])
	}
}

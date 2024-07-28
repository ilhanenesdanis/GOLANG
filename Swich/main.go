package main

func main() {
	foo := -1

	switch {
	case foo == 1:
		println("One")
	case foo == 2:
		println("two")
	case foo > 2:
		println("greater")
	default:
		println("other")
	}
}

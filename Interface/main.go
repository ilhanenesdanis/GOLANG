package main

import (
	"fmt"
	"strconv"
)

func CarExecute(c Carface) {
	fmt.Println()
	fmt.Println("Araç Bilgi : \n" + c.Information())

	fmt.Println()

	msg := ""
	isRun := c.Run()

	if isRun {
		msg = "Araç Çalişiyor"
	} else {
		msg = "Araç Çalişmiyor"
	}
	fmt.Println(msg)

	isStop := c.Stop()

	if isStop {
		msg = "Araç Duruyor"
	} else {
		msg = "Araç Çalışıyor"
	}

	fmt.Println(msg)

}

func main() {

	bmw := new(Bmw)
	bmw.Brand = "BMW"
	bmw.Model = "330Cİ"
	bmw.Color = "Red"
	bmw.Speed = 250
	bmw.Price = 1.5
	bmw.Special = true
	bmw.Run()
	//fmt.Println(bmw.Information())

	CarExecute(bmw)

}

//Interface

type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

// Base Struct
type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

type SpecialProduction struct {
	Special bool
}

// Bmw
type Bmw struct {
	Car //composition
	SpecialProduction
	Rims string
}

func (_ Bmw) Run() bool {
	return true
}
func (_ Bmw) Stop() bool {
	return true
}
func (x *Bmw) Information() string {

	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color: " + x.Color + "\n" + "\t" + "Speed: " + strconv.Itoa(x.Speed) + "\n\t" + "Price: " + strconv.Itoa(int(x.Price))
	add := "Yes"
	if x.Special {
		ret += "\n\t" + "Special: " + add
	}

	return ret
}

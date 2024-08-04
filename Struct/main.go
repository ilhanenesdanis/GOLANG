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

type Vertex struct {
	x int
	y int
}

//User
type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

func NewUser() *User {
	u := new(User)
	u.Pay = NewPayment()
	return u
}

//Payment
type Payment struct {
	Salary float64
	Bonus  float64
}

func NewPayment() *Payment {
	u := new(Payment)

	return u
}

func (u User) GetFullName() string {

	return u.FirstName + " " + u.LastName
}
func (u User) GetUserName() string {
	return u.UserName
}
func (u User) GetPayment() float64 {
	pay := u.Pay.Salary + u.Pay.Bonus
	return pay
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

	fmt.Println(Vertex{1, 2})

	s := new(Vertex)

	s.x = 1
	s.y = 3

	//UserExample

	user1 := &User{
		ID:        1,
		FirstName: "İlhan",
		LastName:  "Daniş",
		Age:       22,
		UserName:  "danisenes1@gmail.com",
		Pay: &Payment{
			Salary: 55400,
			Bonus:  150,
		},
	}

	fmt.Println(user1)
	fmt.Println(user1.Pay)

	fmt.Println(user1.GetFullName())
	fmt.Println(user1.GetUserName())
	fmt.Println(user1.GetPayment())

}

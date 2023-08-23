package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (x customer) save() {
	fmt.Println("Eklendi")
}

func (x customer) update() {
	fmt.Println("Güncellendi")
}

func Demo2() {
	c := customer{firstName: "Onur", lastName: "Bilici", age: 19}
	c.save()
	c.update()
}

package test_struct

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) setName(name string) {
	p.name = name
}

func (p Person) getName() string {
	return p.name
}

func NewPerson(name string, age int) Person {
	return Person{
		name: name,
		age:  age,
	}
}

func Run() {
	fmt.Println("test_struct")
	p1 := Person{
		name: "conan",
		age:  20,
	}
	p2 := NewPerson("conan2", 21)

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.getName())
}

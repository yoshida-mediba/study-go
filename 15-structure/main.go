package main

import (
	"fmt"
	"time"
)

type Point struct {
	X int
	Y int
}

type Person struct {
	id   uint
	name string
	time string
}

// Structure Method
func (p *Person) updatePersonTime() {
	p.time = time.Now().Format("2006-01-02 15:04:05")
}

func NewPerson(id uint, name string) *Person {
	p := new(Person)
	p.id = id
	p.name = name
	p.updatePersonTime()
	return p
}

func main() {
	var pt Point
	pt.X = 10
	pt.Y = 100
	fmt.Println(pt.X, pt.Y)

	pt2 := Point{20, 200}
	fmt.Println(pt2.X, pt2.Y)

	p := &Person {
		id: 21339,
		name: "Hiroki Yoshida",
	}

	p.updatePersonTime()
	fmt.Println(p)

	p2 := NewPerson(21339, "Hiroki Yoshida")
	fmt.Println(p2)
}

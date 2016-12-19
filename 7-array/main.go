package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%#v [%T]\n", a, a)
	fmt.Println("-----")

	var b [5]int
	fmt.Printf("%#v [%T]\n", b, b)
	fmt.Println("-----")

	b = a
	fmt.Printf("%#v [%T]\n", a, a)
	fmt.Printf("%#v [%T]\n", b, b)
	fmt.Println("-----")

	a[0] = 11
	b[1] = 22
	fmt.Printf("%#v [%T]\n", a, a)
	fmt.Printf("%#v [%T]\n", b, b)
	fmt.Println("-----")
}

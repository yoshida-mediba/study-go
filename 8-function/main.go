package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")

	voidFunc()
	fmt.Printf("%T\n", stdFunc())
	a, b := multiFunc()
	fmt.Printf("%T, %T\n", a, b)

	_, r := multiFunc()
	fmt.Printf("%T\n", r)

	result, err := doSomething()
	fmt.Printf("%T\n", result)
	fmt.Printf("%T\n", err)

	f := func(x, y int) int {
		return x + y
	}
	fmt.Printf("%d\n", f(a, b))

	f2 := func() func(z int) int {
		return func(z int) int {
			return z * 2
		}
	}

	fmt.Printf("%d\n", f2()(100))
}

func voidFunc() {
	fmt.Println("voidFunc()")
	return
}

func stdFunc() bool {
	fmt.Println("stdFunc()")
	return true
}

func multiFunc() (int, int) {
	return 10, 10
}

func doSomething() (int, string) {
	return 10, "Something Error Happens"
}

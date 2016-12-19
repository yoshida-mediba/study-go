package main

import (
	"fmt"
)

func main() {
	var i int
	for {
		if (i >= 5) {
			break
		} else {
			fmt.Println("i =", i)
			i++
		}
	}

	for j := 0; j < 5; j++ {
		fmt.Println("j =", j)
	}

	var p int
	for p < 5 {
		fmt.Println("p =", p)
		p++
	}

	if _, err := doSomething(); err != "" {
		fmt.Println(err)
	}

	fruits := [3]string{"Apple", "Banana", "Cherry"}
	for i, s := range fruits {
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}

	n := 10
	switch n {
	case 10:
		fmt.Println("n == 10!")
	default:
		fmt.Println("default!")
	}

	switch {
	case n > 10 || n < 10:
		fmt.Println("n <> 10!")
	default:
		fmt.Println("n == 10!")
	}

	var x interface{} = 3.14
	_, isInt := x.(int)
	_, isFloat64 := x.(float64)
	_, isString := x.(string)
	fmt.Printf("%v, %v, %v\n", isInt, isFloat64, isString)

	if x == nil {
		fmt.Println("x is nil")
	} else if xi, isInt := x.(int); isInt {
		fmt.Printf("x is integer: %d\n", xi)
	} else if xf, isFloat64 := x.(float64); isFloat64 {
		fmt.Printf("x is float: %v\n", xf)
	} else if xs, isString := x.(string); isString {
		fmt.Println(xs)
	}

	switch v := x.(type) {
	case bool:
		fmt.Println("bool", v)
	case string:
		fmt.Println("string", v)
	case float64:
		fmt.Println("float64", v)
	}
}

func doSomething() (int, string) {
	return 10, "ERROR!!"
}

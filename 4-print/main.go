package main

import (
	"fmt"
)

func main() {
	fmt.Println("stdout")
	println("stderr")

	fmt.Printf("%d-%d-%d\n", 2016, 12)
	fmt.Printf("%d-%d-%d", 2016, 12, 19, 15)

	fmt.Println("")

	fmt.Printf("数値=%v, 文字列=%v, 配列=%v\n", 5, "Golang", [...]int{1,2,3})
	fmt.Printf("数値=%#v, 文字列=%#v, 配列=%#v\n", 5, "Golang", [...]int{1,2,3})
	fmt.Printf("数値=%T, 文字列=%T, 配列=%T\n", 5, "Golang", [...]int{1,2,3})
}

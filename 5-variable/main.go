package main

import (
	"fmt"
)

var n3 = 100

func main() {
	var n int
	fmt.Printf("%#v [%T]\n", n, n)
	fmt.Println("-----")

	// 複数
	var x,y,z int
	fmt.Printf("%#v [%T]\n", x, x)
	fmt.Printf("%#v [%T]\n", y, y)
	fmt.Printf("%#v [%T]\n", z, z)
	fmt.Println("-----")

	// 複数
	var (
		p, q int
		name string
	)
	fmt.Printf("%#v [%T]\n", p, p)
	fmt.Printf("%#v [%T]\n", q, q)
	fmt.Printf("%#v [%T]\n", name, name)
	fmt.Println("-----")

	// 暗黙的な定義
	b := true
	i := 1
	f := 3.14
	s := "abc"
	fmt.Printf("%#v [%T]\n", b, b)
	fmt.Printf("%#v [%T]\n", i, i)
	fmt.Printf("%#v [%T]\n", f, f)
	fmt.Printf("%#v [%T]\n", s, s)
	fmt.Println("-----")

	// 代入
	x, y, z = 1, 10, 100
	fmt.Printf("%#v [%T]\n", x, x)
	fmt.Printf("%#v [%T]\n", y, y)
	fmt.Printf("%#v [%T]\n", z, z)
	fmt.Println("-----")

	// 複数var初期値
	var (
		n1 = 1
		n2 = "true"
	)
	fmt.Printf("%#v [%T]\n", n1, n1)
	fmt.Printf("%#v [%T]\n", n2, n2)
	fmt.Println("-----")


	n3 ++
	fmt.Printf("%#v [%T]\n", n3, n3)
	fmt.Println("-----")
}

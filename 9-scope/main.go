package main

import (
	"fmt"
	"./conf"
	. "math"
)

func main() {
	fmt.Println(conf.GetX(), conf.X)
	fmt.Println(conf.GetY()) // can not get conf.y

	fmt.Println(Pi)
}

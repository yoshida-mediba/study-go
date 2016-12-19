package main

import (
	"fmt"
	"math"
)

var sum uint32

func main() {
	fmt.Printf("uint32 max value= %d\n", math.MaxUint32)

	fmt.Printf("%#v [%T]\n", sum, sum)
	if ! addSum(10) {
		println("overflow")
	}
	fmt.Printf("%#v [%T]\n", sum, sum)
	if ! addSum(4294967295) {
		println("overflow")
	}
	fmt.Printf("%#v [%T]\n", sum, sum)
	fmt.Println("-----")

	// overflow
	sum += 4294967295
	fmt.Printf("%#v [%T]\n", sum, sum)
	sum += 4294967295
	fmt.Printf("%#v [%T]\n", sum, sum)
}

func addSum(b uint32) bool {
	if (math.MaxUint32 - b) < sum {
		return false
	} else {
		sum += b
		return true
	}
}

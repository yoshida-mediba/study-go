package main

import (
	"fmt"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	doPanic()
}

func doPanic() {
	panic("ERROR!!")
}

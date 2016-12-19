package main

import (
	"fmt"
	"time"
)

var sleep time.Duration
var S = "A"

func init() {
	sleep = 1 * time.Millisecond
	S += "B"
}

func main() {
	go sub()
	time.Sleep(sleep)
	S += "C"
	fmt.Println(S)
}

func sub() {
	for {
		fmt.Println("SUB!!")
	}
}

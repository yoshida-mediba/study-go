package main

import (
	"fmt"
	"strconv"
)

func main() {
	m := make(map[int]int)
	m[1] = 10
	m[2] = 20
	for _, v := range m {
		fmt.Println(v)
	}

	m2 := map[int]string {
		1: "Taro",
		2: "Hanako",
		3: "Jiro",
	}
	for _, v := range m2 {
		fmt.Println(v)
	}

	m3 := map[string]map[string]string {
		"Monkey": {
			"Eats": "Banana",
			"Live": "Jangle",
		},
		"Elephant": {
			"Eats": "Apple",
			"Live": "Africa",
		},
	}
	fmt.Println("Length: "+strconv.Itoa(len(m3)))

	m3["Human"] = map[string]string {
		"Eats": "Bread",
		"Live": "Earth",
	}
	fmt.Println("Length: "+strconv.Itoa(len(m3)))

	for k, v := range m3 {
		fmt.Println(k + " {")
		for k2, v2 := range v {
			fmt.Println(" ", k2, v2)
		}
		fmt.Println("}")
	}
}

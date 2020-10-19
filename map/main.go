package main

import (
	"fmt"
)

func main() {
	singleMap()
	mutipleMap()
}

func singleMap() {
	fmt.Println("singleMap ........")
	var numbers = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(numbers)
	fmt.Println(numbers["two"])

	var colors = make(map[string]string)
	colors["red"] = "#f00"
	colors["green"] = "#0f0"
	colors["blue"] = "#00f"

	fmt.Println(colors)
	fmt.Println(colors["green"])
}

func mutipleMap() {
	fmt.Println("mutipleMap ........")
	var courses = make(map[string]map[string]int)
	courses["Java"] = make(map[string]int)
	courses["Java"]["price"] = 1200
	courses["Java"]["code"] = 100

	courses["Golang"] = make(map[string]int)
	courses["Golang"]["price"] = 1000
	courses["Golang"]["code"] = 200

	courses["NodeJS"] = make(map[string]int)
	courses["NodeJS"]["price"] = 900
	courses["NodeJS"]["code"] = 200

	fmt.Println("courses all       : ", courses)
	fmt.Println("courses Java code : ", courses["Java"]["code"])
}

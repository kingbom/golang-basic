package main

import (
	"fmt"
)

func main() {
	var numbers []int = []int{1, 2, 3, 4}
	var numbers2 = []int{1, 2, 3, 4}
	var skills [4]string
	skills[0], skills[1], skills[2] = "java", "node", "golang"
	skills[3] = "devOps"

	for _, item := range numbers {
		fmt.Print(item, ",")
	}
	fmt.Println("")

	for _, item := range numbers2 {
		fmt.Print(item, ",")
	}
	fmt.Println("")

	for _, item := range skills {
		fmt.Println("skill = ", item)
	}

}

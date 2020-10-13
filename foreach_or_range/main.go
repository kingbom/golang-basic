package main

import (
	"fmt"
)

func main() {
	skills := []string{"java", "node", "golang"}

	for i, skill := range skills {
		fmt.Printf("%d. %s \n", i, skill)
	}

	for _, skill := range skills {
		fmt.Printf("%s \n", skill)
	}
}

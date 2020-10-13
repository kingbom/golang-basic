package main

import (
	"fmt"
	"time"
)

func main() {
	basicIf()
	sortIf()
	switchCase()
}

func basicIf() {
	num := 10
	if num == 10 {
		fmt.Println("is 10")
	} else {
		fmt.Println("not  10")
	}

	if num >= 10 || num < 20 {
		fmt.Println("is num >= 10 || num < 20")
	} else {
		fmt.Println("not num >= 10 || num < 20")
	}
}

func sortIf() {
	if result := doSometing(); result == "ok" {
		fmt.Println("is ok")
	} else {
		fmt.Println("not ok")
	}
}

func switchCase() {
	index := 2
	switch index {
	case 0:
		fmt.Println("index is 0")
		break
	case 1:
		fmt.Println("index is 1")
		break
	case 2:
		fmt.Println("index is 2")
		break
	default:
		fmt.Println("index is 2")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func doSometing() string {
	return "ok"
}

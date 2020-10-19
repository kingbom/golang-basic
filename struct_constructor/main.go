package main

import (
	"oop/model"
)

func main() {
	e := model.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}

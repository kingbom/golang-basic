package main

import (
	"oop/model"
)

func main() {
	e := model.Init("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()

	e = model.Init("Bom", "Adolf", 30, 20)
	e.LeavesRemaining()
}

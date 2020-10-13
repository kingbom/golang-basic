package main

import (
	"fmt"
)

func main() {
	msg := "old message"
	var msgPointer *string = &msg
	fmt.Printf("msg        before    is %s \n", msg)
	fmt.Printf("msg assigned pointer is %s \n", *msgPointer)

	changeMessage(&msg, "new message 1")
	fmt.Printf("msg after change pointer is %s \n", msg)

	changeMessage(msgPointer, "new message 2")
	fmt.Printf("msg after change pointer is %s \n", msg)
}

func changeMessage(aPointer *string, newMessage string) {
	*aPointer = newMessage
}

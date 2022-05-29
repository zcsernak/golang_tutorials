package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)

	var msg = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye"
	// the go runtime won't interrupt the main thread until the sleep call
	// so the anyonymous function executed in a new goroutine will see "Goodbye" in the msg variable
	time.Sleep(100 * time.Millisecond)

	var msg2 = "Hello"
	go func(msg string) {
		fmt.Println(msg)
	}(msg2)

	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}

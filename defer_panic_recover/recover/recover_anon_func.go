package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	// anonymous function
	defer func() {
		// recover() returns the error which is causing the application to be panicking
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("end")
}

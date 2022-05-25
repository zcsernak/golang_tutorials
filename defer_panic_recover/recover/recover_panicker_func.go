package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			// can't handle panic: rethrowing it
			panic(fmt.Sprintf("from panicker function: %s", err))
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}

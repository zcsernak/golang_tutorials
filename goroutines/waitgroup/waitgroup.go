package main

import (
	"fmt"
	"sync"
)

// WaitGroup is designed to be used globally like this, it is safe
var wg = sync.WaitGroup{}

func main() {
	var msg = "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	wg.Wait()
}

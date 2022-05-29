package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	// go by default gives you the number of operating system threads equal to the number of cores in the machine
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(100)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

package main

import (
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	incrementerExample()
}

type Writer interface {
	Write([]byte) (int, error)
}

// ConsoleWriter implements the Writer interface
// by having a method with the same signiture as defined in the interface:
// Write([]byte) (int, error)
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func incrementerExample() {
	myInt := IntCounter(0) // casting int to IntCounter
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

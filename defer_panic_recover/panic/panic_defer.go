package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened") // panicks happen after the deffered statements were executed
	fmt.Println("end")
}

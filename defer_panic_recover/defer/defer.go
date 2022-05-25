package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	deferExample()
	deferLifo()
	resourceRequest()
	deferVariableExample()
}

func deferExample() {
	// deferred functions are executed after the function finishes its last statement, but before it returns
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}

func deferLifo() {
	// deferred functions are executed in LIFO order (last in first out)
	// defer is usually used for closing resources and it makes sense to close them in the opposite order as they were opened,
	// because they might depend on another one
	defer fmt.Println("defelerLifo: start")
	defer fmt.Println("defelerLifo: middle")
	defer fmt.Println("defelerLifo: end")
}

func resourceRequest() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err, res)
	}
	defer res.Body.Close()
	// in for loops you might want to close resources without defer,
	// because all of the resources will stay open until the function finishes running eating up a lot of memory
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(robots[0:13]))
}

func deferVariableExample() {
	a := "start"
	// function calls take the arguments at the time defer is called
	// not at the time the function is executed
	defer fmt.Println(a)
	a = "end"
}

package main

import (
	"fmt"
	"strconv"
)

var l int = 42

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

var (
	counter = 1
)

var PublicExportedInteger int = 5
var notExported string = "use camelCase"

var k int = 77

func main() {
	var i int
	i = 42
	//var j float32 = 27
	fmt.Printf("%v, %T\n", k, k)
	k := 99.2
	fmt.Printf("%v, %T\n", k, k)

	var m float32
	m = float32(i)
	fmt.Printf("%v, %T\n", m, m)

	n := int(k)
	fmt.Printf("%v, %T\n", n, n)

	var s string
	s = string(i) // string of bytes
	fmt.Printf("%v, %T\n", s, s)

	s = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", s, s)

}

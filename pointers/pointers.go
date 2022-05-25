package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(&a, b)
	a = 27
	fmt.Println(a, *b)
	*b = 69
	fmt.Println(a, *b)

	c := 77
	swap(&a, &c)
	fmt.Println(a, c)

	fmt.Println()
	pointerArithmetic()
	fmt.Println()
	structExample()
	fmt.Println()
	newFunctionExample()
	fmt.Println()
	sliceExample()
	fmt.Println()
	mapExample()
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func pointerArithmetic() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1] /* - 4  cannot convert 4 (untyped int constant) to *int */
	fmt.Printf("%v %p %p\n", a, b, c)

	// unsafe package: https://pkg.go.dev/unsafe
}

type myStruct struct {
	foo int
}

func structExample() {
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)
}

func newFunctionExample() {
	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)
	fmt.Println(ms)
	(*ms).foo = 42
	fmt.Println((*ms).foo)
	ms.foo = 88 // syntactic sugar
	fmt.Println(ms.foo)
}

func sliceExample() {
	a := []int{1, 2, 3}
	b := a // slices have a pointer to an array, so this statement copies the pointer, and not the underlying data itself
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)
}

func mapExample() {
	a := map[string]string{"foo": "bar", "baz": "buz"}
	b := a
	fmt.Println(a, b)
	a["foo"] = "qux"
	fmt.Println(a, b)
}

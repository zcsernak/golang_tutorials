package main

import (
	"fmt"
)

func main() {
	// Array
	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	students := [...]string{"Joe", "Jim", "Jane"}
	fmt.Printf("Students: %v\n", students)

	students[0] = "Johnny"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of Students: %v\n", len(students))
	fmt.Printf("Capacity of Students array: %v\n", cap(students))

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	a := [...]int{1, 2, 3}
	b := a  // array being copied
	c := &a // pointer
	b[1] = 5
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	c[2] = 55
	fmt.Println("a =", a)

	// Slice
	aSlice := []int{1, 2, 3}
	fmt.Println("aSlice", aSlice)
	fmt.Println("aSlice len", len(aSlice))
	fmt.Println("aSlice cap", cap(aSlice))
	aSlice = append(aSlice, 4, 5)
	fmt.Println("aSlice len", len(aSlice))
	fmt.Println("aSlice cap", cap(aSlice))

	bSlice := aSlice // they are backed by the same array
	bSlice[1] = 88
	fmt.Println("aSlice =", aSlice)
	fmt.Println("bSlice =", bSlice)

	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	y := x[:]   // slice of all elements
	z := x[3:]  // inclusive
	zs := x[:6] // exclusive
	zsé := x[7:9]
	y[0] = 444
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(zs)
	fmt.Println(zsé)

	madeSlice := make([]int, 3, 5)
	fmt.Printf("%v, %T\n", madeSlice, madeSlice)
	fmt.Printf("Length: %v\n", len(madeSlice))
	fmt.Printf("Capacity: %v\n", cap(madeSlice))

	madeSlice = append(madeSlice, []int{2, 3, 4, 5}...)
	fmt.Printf("Length: %v\n", len(madeSlice))
	fmt.Printf("Capacity: %v\n", cap(madeSlice))

	stack1 := []int{0, 1, 2, 3, 4}
	stack2 := stack1[1:] // pop first element
	fmt.Println(stack2)
	stack2 = stack1[:len(stack1)-1] // remove last element
	fmt.Println(stack2)
	stack2 = append(stack1[:2], stack1[3:]...) // remove element from middle
	fmt.Println(stack2)
	fmt.Println(stack1)

}

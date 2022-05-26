package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Go!")

	fmt.Println()
	sayMessage("Hi!")
	for i := 0; i < 5; i++ {
		sayMessageWithIdx("Hello Go!", i)
	}
	fmt.Println()
	greetName("Hello", "Csocsó")

	fmt.Println()
	greeting := "Hello"
	name := "Geralt"
	sayGreeting(&greeting, &name)
	fmt.Println(greeting, name)

	fmt.Println()
	s1 := sum("The sum is", 1, 2, 3, 4, 5)
	fmt.Println("Yup the sum is", s1)

	fmt.Println()
	s2 := sumPtr(1, 2, 3, 4, 5)
	fmt.Println("SumPtr", *s2)

	fmt.Println()
	namedReturnValue(1, 2, 3, 4, 5)
	fmt.Println("Yup the sum is still", s1)

	fmt.Println()
	d, err := multipleReturnValues(5.0, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}

	fmt.Println()
	func() {
		anonMsg := "I am an anonymous function"
		fmt.Println(anonMsg)
	}()

	fmt.Println()
	for i := 0; i < 5; i++ {
		func(i int) { // if this would be running asynchronously then i could be incemented multiple times and exceeding the exit condition i < 5
			fmt.Println(i)
		}(i)
	}

	fmt.Println()
	// var f func() =
	f := func() {
		fmt.Println("Press f to pay respects")
	}
	f()

	fmt.Println()
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, fmt.Errorf("0-val ne ossz, kössz")
		} else {
			return a / b, nil
		}
	}
	d2, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d2)
	}

	fmt.Println()
	// Method
	g := greeter{
		greeting: "Csöváz",
		name:     "Peepoo",
	}
	// Method is a function executed in a known context
	// known context = any type
	g.greet()
	fmt.Println("The new name is:", g.name)
	g.greetAndEmptyName()
	fmt.Println("The new name is:", g.name)
}

func sayMessage(msg string) {
	fmt.Println(msg)
}

func sayMessageWithIdx(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

func greetName(greeting, name string) {
	fmt.Println(greeting, name)
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

// only one variatic parameter is allowed, and it has to be the last one
func sum(msg string, values ...int) int {
	fmt.Printf("%v, %T\n", values, values) // type = slice
	result := 0
	for _, num := range values {
		result += num
	}
	fmt.Println(msg, result)
	return result
}

func sumPtr(values ...int) *int {
	result := 0
	for _, num := range values {
		result += num
	}
	return &result // promoted to shared (a.k.a. heap) memory from the local stack memory
}

func namedReturnValue(values ...int) (result int) {
	fmt.Println(values)
	for _, num := range values {
		result += num
	}
	return
}

func multipleReturnValues(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("gyökér, ne osszámán nullával")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

// g = value receiver
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}

// g = pointer receiver
func (g *greeter) greetAndEmptyName() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}

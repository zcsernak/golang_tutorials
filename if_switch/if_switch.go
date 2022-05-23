package main

import (
	"fmt"
	"math"
)

func main() {
	if true {
		fmt.Println("The test is true")
	}

	statePopulations := map[string]int{
		"California":  39250017,
		"Texas":       27862596,
		"Florida":     20612439,
		"New York":    19745289,
		"Pensylvania": 12802503,
		"Illinois":    12801539,
		"Ohio":        11614373,
	}

	// initializer syntax
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop) // pop is block scoped
	}

	number := 50
	guess := 30

	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100")
	}
	if guess >= 0 && guess <= 100 {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
	}

	fmt.Println(number <= guess, number >= guess, number != guess)

	fmt.Println(!true)

	// short circuit
	if guess == 30 || returnTrue() {
		fmt.Println("Your guess is 30")
	}

	elseIf(guess, number)
	floatTest()
	switchExample()
	switchExample2()
	taglessSwitch()
	typeSwitch()
	breakSwitch()
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}

func elseIf(guess int, number int) {
	if guess < 1 {
		fmt.Println("The guess must be greater than 1")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
	}
}

func floatTest() {
	myNum := 0.123
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are in the same margin")
	} else {
		fmt.Println("These are different")
	}
}

func switchExample() {
	switch 4 {
	case 1:
		fmt.Println("one")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}
}

func switchExample2() {
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, fice or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}
}

func taglessSwitch() {
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // other logic won't be executed
	case i <= 20: // overlapping is allowed here
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}
}

func typeSwitch() {
	var i interface{} = [3]int{}
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is an float64")
	case string:
		fmt.Println("i is a string")
	case [3]int:
		fmt.Println("i is [3]int")
	case [2]int:
		fmt.Println("i is [2]int")
	default:
		fmt.Printf("i is another type: %T\n", i)
	}

}

func breakSwitch() {
	var i interface{} = []int{}
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		break
		fmt.Println("This won't print")
	case float64:
		fmt.Println("i is an float64")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Printf("i is another type: %T\n", i)
	}
}

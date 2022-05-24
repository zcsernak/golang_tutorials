package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
	}

	fmt.Println()
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()
	// while
	i = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println()
	// infinite loop
	for {
		fmt.Println("Hello")
		if i == 7 {
			break
		} else {
			i++
		}
	}

	fmt.Println()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println()
	// label
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	fmt.Println()
	// for range loop
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println("index:", k, "value:", v)
	}

	fmt.Println()
	statePopulations := map[string]int{
		"California":  39250017,
		"Texas":       27862596,
		"Florida":     20612439,
		"New York":    19745289,
		"Pensylvania": 12802503,
		"Illinois":    12801539,
		"Ohio":        11614373,
	}

	fmt.Println()
	for k, v := range statePopulations {
		fmt.Println("key:", k, "value:", v)
	}

	fmt.Println("\nOnly keys")
	for k := range statePopulations {
		fmt.Println(k)
	}

	fmt.Println("\nOnly values")
	for _, v := range statePopulations {
		fmt.Println(v)
	}

	fmt.Println()
	str := "Hello Go!"
	for k, v := range str {
		fmt.Println("index:", k, "unicode:", v, "string:", string(v))
	}

}

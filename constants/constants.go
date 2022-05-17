package main

import (
	"fmt"
)

const a = 27

const (
	x = iota // ~counter starting from 0
	y
	z // iota lesz ez is
)

const (
	x2 = iota
)

const (
	_ = iota + 5 // errorSpecialist, ignore first value by assigning to blank identifier a.k.a don't care
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	//const myPi float64 = math.Sin(1.57) // math.Sin(1.57) (value of type float64) is not constant <- Value must be calculable at compile time

	// Arrays are always mutable, you can't create a constant array

	const a int = 14 // shadowing the "a" variable from outer scope
	fmt.Printf("%v, %T\n", a, a)

	b := 7
	fmt.Printf("%v, %T\n", a+b, a+b)

	const converted = 69
	var d int16 = 1
	fmt.Printf("%v, %T\n", converted+d, converted+d) // compiler replaces the literal 'converted' with 69 -> implicit conversion is possible

	// enumerated constants
	fmt.Printf("%v, %T\n", x, x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(x2) // iota is constant block scoped

	specialistType := dogSpecialist
	fmt.Printf("%v\n", specialistType == dogSpecialist)

	var notSpecializedSpecialist int
	fmt.Printf("%v\n", notSpecializedSpecialist == catSpecialist)
	fmt.Println(catSpecialist)

	fmt.Println("KiloByte =", KB)
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	var roles byte = isAdmin | canSeeAfrica | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)

}

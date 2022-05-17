package main

import (
	"fmt"
	"strconv"
)

func main() {
	// bool
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	m := 1 == 1
	o := 1 == 2

	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", o, o)

	var p bool // zero value = false
	fmt.Printf("%v, %T\n", p, p)

	// numeric
	// zero value = 0

	// int, int8, int16, int32, int64
	q := 42
	fmt.Printf("%v, %T\n", q, q)

	// uint, uint8, uint16, uint32
	var s uint16 = 42
	fmt.Printf("%v, %T\n", s, s)

	a := 10 // 1010
	b := 3  // 0011
	fmt.Println("\nNumber operators")
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	var c int8 = 4
	fmt.Println(a + int(c))

	fmt.Println("\nBit operators")
	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 1000
	fmt.Println(strconv.FormatInt(int64(a)&^int64(b), 2))

	// bit shifting
	d := 8              // 2^3
	fmt.Println(d << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(d >> 3) // 2^3 / 2^3 = 2^0

	// floating point numbers
	// float32, float64
	fmt.Println("\nFloats")
	e := 3.14
	e = 2.1e14
	fmt.Printf("%v, %T\n", e, e)

	f := 10.2
	g := 3.7
	fmt.Println(f / g)

	// complex number
	// complex64, complex128
	fmt.Println("\nComplex")
	var z complex128 = 1.5 + 2i
	fmt.Printf("%v, %T\n", z, z)
	fmt.Printf("%v, %T\n", real(z), real(z))
	fmt.Printf("%v, %T\n", imag(z), imag(z))

	var y complex64 = complex(5, 12)
	fmt.Printf("%v, %T\n", y, y)

	fmt.Println("\nStrings") // any utf8
	s1 := "this is a string"
	fmt.Printf("%v, %T\n", s1, s1)
	fmt.Printf("%v, %T\n", string(s1[2]), string(s1[2]))
	// strings are immutable
	// str[2] = "u"

	s2 := "this is a string too"
	fmt.Printf("%v, %T\n", s1+s2, s1+s2)

	s3 := []byte(s1)
	fmt.Printf("%v, %T\n", s3, s3)

	fmt.Println("\nRune")     // any utf32
	var roney_rune rune = 'a' // int32 alias
	fmt.Printf("%v, %T\n", roney_rune, roney_rune)
}

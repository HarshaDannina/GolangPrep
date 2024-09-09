package main

import "fmt"

func main() {

	fmt.Println("hello world")

	// ****************
	// CHAPTER-1: Variables
	// ****************

	// Boolean
	var a bool
	// Whole Numbers
	var b int
	var c int8
	var d int16
	var e int32
	var f int64
	// Positive Numbers
	var g uint
	var h uint8
	var i uint16
	var j uint32
	var k uint64
	var l byte
	var m rune
	// Decimal Numbers
	var n float32
	var o float64
	// Imaginary Numbers
	var p complex64
	var q complex128
	// String
	var r string

	//Formating variables
	// %v - Default representation
	// %s - String
	// %d - Integer in Decimal
	// %f - Decimal
	// %q - String

	fmt.Printf("%v %v %v %v %v %v %v %v %v %v %v %v %v %f %f %f %f %q\n", a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)

	// Short Variable Declaration
	s := "hello"
	t := 2
	u := 3.14
	var v = true

	fmt.Printf("%T %T %T %T\n", s, t, u, v)

	// Type Casting
	t = int(u)
	u = float64(t)

	fmt.Printf("%v %f\n", t, u)

	// Default Types
	// bool string int uint32 byte rune float64 complex128

	// Constants
	const planA = "a"
	const pi = 3.14
	const radius = 10
	const area = pi * radius * radius

	fmt.Println(area)
	var score float64 = 89.9999999999
	msg := fmt.Sprintf("Hi %s, Congrats!, Your score is %.3f", "Praveen", score)
	fmt.Println(msg)

	// Conditional Operators
	// == equals to
	// != not equal to
	// < less than
	// > greater than

	if score >= 90 {
		fmt.Println("Selected")
	} else if score > 60 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}

	if percentage := score * 10; percentage > 900 {
		fmt.Println("Selected")
	} else if percentage > 600 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}

	// Functions
	addition := add(1, 2)
	subtraction := sub(2, 1)
	multiplication, division, err := muldiv(1, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v %v %v %f\n", addition, subtraction, multiplication, division)

	// Struct
	myCar := Car{
		Color: "black",
		Engine: Engine{ // Embeded Struct
			Stroke:   4,
			Cylender: "v8",
		},
		BackWheel: Wheel{},
	}

	myCar.FrontWheel.Radius = 22
	myCar.FrontWheel.Material = "rubber"

	myCar.BackWheel = Wheel{
		Radius:   34,
		Material: "rubber",
	}

	fmt.Println(myCar)
	fmt.Println(myCar.Stroke)

}

// ****************
// CHAPTER-2: Functions
// ****************

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func muldiv(x, y int) (mul int, div float64, err error) {
	if y == 0 {
		return x * y, 0, fmt.Errorf("Can not be divisible by 0")
	}
	return x * y, float64(x) / float64(y), nil
}

// ****************
// CHAPTER-3: Struct
// ****************

type Car struct {
	Color      string
	Engine              // Embeded Struct
	FrontWheel struct { // Nested Struct
		Radius   int
		Material string
	}
	BackWheel Wheel
}

type Engine struct {
	Stroke   int
	Cylender string
}

type Wheel struct {
	Radius   int
	Material string
}

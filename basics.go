package main

import "fmt"

func main() {

	fmt.Println("hello world")

	// ****************
	// CHAPTER-1: Variables
	// ****************
	fmt.Println("### CH-1: Variables")
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
	str := fmt.Sprintf("Hi %s, Congrats!, Your score is %.3f", "Praveen", score)
	fmt.Println(str)

	// Conditional Operators
	// == equals to
	// != not equal to
	// < less than
	// > greater than
	fmt.Println("## Conditional Operators")
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
	fmt.Println("### CH-2: Functions")
	addition := add(1, 2)
	subtraction := sub(2, 1)
	multiplication, division, err := muldiv(1, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v %v %v %f\n", addition, subtraction, multiplication, division)

	// Struct
	fmt.Println("### CH-3: Structs")
	myCar := Car{
		Color: "black",
		Engine: Engine{ // Embeded Struct
			Stroke:   4,
			Cylinder: "v8",
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
	// Struct method
	fmt.Println("## Struct Method")
	fmt.Println(myCar.printInfo())

	// Interfaces
	fmt.Println("### CH-4: Interface")
	post := postNotification{
		postId: 1,
		userId: 2,
	}

	msg := messageNotification{
		senderId: 1,
		message:  "Hello",
	}

	sendNotification(post)
	sendNotification(msg)

	// Multiple Interfaces
	fmt.Println("## Multiple Interfaces")
	openApp(post, post)
	openApp(msg, msg)

	// Type Assertions
	fmt.Println("## Type Assertions")
	printNotification(post)
	printNotification(msg)

	fmt.Println("## Type Switches")
	checkNotification(post)
	checkNotification(msg)
	checkNotification(nil)
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
	Cylinder string
}

type Wheel struct {
	Radius   int
	Material string
}

// Struct Methods
func (car Car) printInfo() string {
	return fmt.Sprintf("{Color: %s, Engine: {Stroke: %v, Cylinder: %s}, Front Wheel: {Radius: %v, Material: %s}, Back Wheel: {Radius: %v, Material: %s}}",
		car.Color, car.Stroke, car.Cylinder, car.FrontWheel.Radius, car.FrontWheel.Material, car.BackWheel.Radius, car.BackWheel.Material)
}

// ****************
// CHAPTER-4: Interface
// ****************

type notification interface {
	pushNotification() string
}

type postNotification struct {
	postId int
	userId int
}

type messageNotification struct {
	senderId int
	message  string
}

func (notify postNotification) pushNotification() string {
	return fmt.Sprintf("%v liked your post %v", notify.userId, notify.postId)
}

func (notify messageNotification) pushNotification() string {
	return fmt.Sprintf("%v sent a message: %s", notify.senderId, notify.message)
}

func sendNotification(notify notification) {
	fmt.Println(notify.pushNotification())
}

// Multiple interfaces
type clickHandler interface {
	onClick() string
}

func (notify postNotification) onClick() string {
	return fmt.Sprintf("Open post %v", notify.postId)
}

func (notify messageNotification) onClick() string {
	return fmt.Sprintf("Open chat with user %v", notify.senderId)
}

func openApp(n notification, c clickHandler) {
	fmt.Println("Notification -> ", n.pushNotification())
	fmt.Println("OnClick -> ", c.onClick())
}

// Type Assertions
func printNotification(n notification) {
	p, ok := n.(postNotification)
	if ok {
		fmt.Println(p.pushNotification())
	}
	m, ok := n.(messageNotification)
	if ok {
		fmt.Println(m.pushNotification())
	}
}

// Type Switches
func checkNotification(n notification) {
	switch s := n.(type) {
	case postNotification:
		fmt.Println("You got a like from user:", s.userId)
	case messageNotification:
		fmt.Println("You got a message from user:", s.senderId)
	default:
		fmt.Println("You didnot received any notification")
	}
}

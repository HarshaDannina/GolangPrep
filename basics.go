package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("hello world")

	// ****************
	// CHAPTER-1: Variables
	// ****************
	fmt.Println("**************** CH-1: Variables ****************")
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
	fmt.Println("**************** CH-2: Functions ****************")
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
	fmt.Println("**************** CH-4: Interface ****************")
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

	// Error Handling
	fmt.Println("**************** CH-5: Error Handling ****************")
	costOfMessage, err := messengerCost("Hi", "Hello")
	fmt.Println(costOfMessage, err)

	costOfMessage, err = messengerCost("Hi, where are you staying", "Mysore")
	fmt.Println(costOfMessage, err)

	costOfMessage, err = messengerCost("Hi, where are you staying", "I am staying in Mysore")
	fmt.Println(costOfMessage, err)

	// Error Interface
	fmt.Println("## Error Interface")
	quo, err := divide(10.0, 0.0)
	fmt.Println(quo, err)
	// errors package
	quo, err = divideStandard(10.0, 0.0)
	fmt.Println(quo, err)

	// Pointers
	fmt.Println("**************** CH-6: Pointers ****************")
	str1 := "Hello"
	str1Ptr := &str1
	fmt.Println(str1Ptr)
	str1 = "World"
	var str1Ptr2 *string = &str1
	fmt.Println(str1Ptr2)

	arr := [...]int{1, 2, 3, 4, 5}
	arr1Ptr := &arr
	fmt.Println(arr1Ptr[0])

	num1 := 5
	fmt.Println(num1)
	changeNumber(&num1)
	fmt.Println(num1)

	// Loops
	fmt.Println("**************** CH-7: Loops ****************")
	loop1(5)
	loop2(5)
	loop3(5)
	loop4(5)
	loop5(5)

	// Arrays / Slices
	fmt.Println("**************** CH-8: Arrays / Slices ****************")
	arr1 := generateArray(4)
	fmt.Println(arr1)
	fmt.Println("## Slices")
	slice1, slice2, slice3 := generateSlice(5, 10)
	fmt.Println(slice1, slice2, slice3)

	slicedArray, err := sliceArray(4)
	fmt.Println(slicedArray, err)
	slicedArray, err = sliceArray(6)
	fmt.Println(slicedArray, err)

	fmt.Println(addingToSlice([]int{1, 2, 3}, 4, 5, 6))

	fmt.Println("## Variadic")
	fmt.Println(variadicSum(1, 2, 3, 4, 5))

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

// ****************
// CHAPTER-5: Error Handling
// ****************

func sendSMS(message string) (float64, error) {
	const maxLen = 10
	const costPerChar = 0.01
	if len(message) > maxLen {
		return 0.0, fmt.Errorf("Failed, your text is over %v characters", len(message)-maxLen)
	}
	return float64(len(message)) * costPerChar, nil
}

func messengerCost(msgFromA, msgFromB string) (float64, error) {
	costOfA, errA := sendSMS(msgFromA)
	costOfB, errB := sendSMS(msgFromB)
	if errA != nil && errB != nil {
		return 0.0, fmt.Errorf("Message from A %s; Message from B %s", errA.Error(), errB.Error())
	} else if errA != nil {
		return costOfB, fmt.Errorf("Message from A %s", errA)
	} else if errB != nil {
		return costOfA, fmt.Errorf("Message from B %s", errB)
	} else {
		return costOfA + costOfB, nil
	}
}

// Error Interface
type NotDividedByZeroError struct {
	dividend float64
}

func (e NotDividedByZeroError) Error() string {
	return fmt.Sprintf("Cannot divide %v by 0", e.dividend)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, NotDividedByZeroError{dividend: a}
	}
	return a / b, nil
}

// Using errors Package
var standardError error = errors.New("This is a standard Zero Division error")

func divideStandard(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, standardError
	}
	return a / b, nil
}

// ****************
// CHAPTER-6: Pointers
// ****************

func changeNumber(num *int) {
	*num = 10
}

// ****************
// CHAPTER-7: Loops
// ****************

func loop1(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

func loop2(n int) {
	for i, j := 0, 1; i < n && j < n; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
}

func loop3(budget float64) {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > budget {
			fmt.Printf("Possible number of messages to send: %v\n", i)
			return
		}
	}
}

func loop4(n int) {
	for n > 0 {
		fmt.Println(n)
		n--
	}
}

func loop5(n int) {
	n = n + 1
	for {
		n--
		if n == 0 {
			break
		}
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// ****************
// CHAPTER-8: Arrays/Slices
// ****************

func generateArray(n int) [5]int {
	var arr1 [5]int
	n = n + 1
	for i := 0; i < n; i++ {
		arr1[i] = i
	}
	return arr1
}

func generateSlice(length, capacity int) (slice1, slice2, slice3 []int) {
	slice1 = make([]int, length, capacity)
	slice2 = make([]int, length)
	slice3 = []int{1, 2, 3, 4, 5}
	return
}
func sliceArray(n int) ([]int, error) {
	arr := [5]int{1, 2, 3, 4, 5}
	if n < len(arr) {
		return arr[:n], nil
	}
	return nil, errors.New("Index Out of bounds")
}

func variadicSum(nums ...int) (sum int) {
	sum = 0
	for _, v := range nums {
		sum += v
	}
	return
}

func addingToSlice(slice1 []int, nums ...int) []int {
	slice1 = append(slice1, nums[0])
	if len(nums) > 2 {
		slice1 = append(slice1, nums[1], nums[2])
	}
	if len(nums) > 3 {
		slice1 = append(slice1, nums[3:]...)
	}
	return slice1
}

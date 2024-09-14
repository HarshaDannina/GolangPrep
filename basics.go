package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
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
	str := fmt.Sprintf("Hi, Congrats!, Your score is %.3f", score)
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

	fmt.Println("## Append")
	fmt.Println(appendCostPerDay([]costs{{0, 4.0}, {1, 2.1}, {1, 3.1}, {5, 2.5}}))

	fmt.Println(createMatrix(5))

	// Maps
	fmt.Println("**************** CH-9: Maps ****************")
	occr := countOccurances("This is a test")
	fmt.Println(occr)
	deleteLeastOccurance(occr)
	fmt.Println(occr)

	// Advance Functions
	fmt.Println("**************** CH-10: Advance Functions ****************")
	msgs := []string{"This is regarding your work", "This is regarding your leave", "This is regarding your home"}
	fmt.Println(messageFormatter(msgs, formatString))

	// Currying
	fmt.Println("## Currying")
	square := doMath(multiplyNumbers)
	double := doMath(addNumbers)

	fmt.Println(square(5))
	fmt.Println(double(5))

	// defer
	fmt.Println("## Defer")
	users := map[string]User{
		"John": {name: "John", admin: false},
		"Kate": {name: "Kate", admin: true},
		"Ken":  {name: "Ken", admin: false},
	}
	fmt.Println(deleteUser(users, "John"))
	fmt.Println(deleteUser(users, "Kate"))

	//clouser
	fmt.Println("## Clousers")
	builder := stringBuilder()
	builder("This")
	builder("is")
	builder("a")
	builder("broken")
	builder("message")
	fmt.Println(builder(""))

	// Anonymous Functions
	fmt.Println("## Anonymous Functions")
	fmt.Println(calculateCostOfMessages([]string{"Hi", "How are you", "Thank you"}, func(s string) float64 {
		cost := 0.0
		for i := 0; i < len(s); i++ {
			cost += 0.01
		}
		return cost
	}))

	// Concurrency
	fmt.Println("**************** CH-11: Concurrency / Channel ****************")
	sendEmail("Hello")
	sendEmail("Hi")
	sendEmail("Thanks")

	// Channel
	fmt.Println("## Channel")
	sendEmailUsingChannel("Hello")
	isEmailSent := <-emailSentChannel // to capture channel output
	fmt.Println(isEmailSent)
	sendEmailUsingChannel("Hi")
	<-emailSentChannel // when channel output was not needed
	sendEmailUsingChannel("Thanks")
	<-emailSentChannel
	close(emailSentChannel) // closing a channel
	printFibonacci(10)      // channel using range
	messageChannel := make(chan string)
	emailChannel := make(chan string)
	createNotification(messageChannel, emailChannel)
	showNotifications(messageChannel, emailChannel)

	// Mutexes
	fmt.Println("**************** CH-12: Mutexes ****************")
	doTransactions(userBalance{balance: 50.0, mux: &sync.RWMutex{}})

	// Generic
	fmt.Println("**************** CH-13: Generic ****************")
	fmt.Println(splitAnySlice([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(splitAnySlice([]float64{1, 2, 3, 4, 5}, 4))
	fmt.Println(splitAnySlice([]string{"1", "2", "3", "4", "5"}, 5))

	fmt.Printf("%v\n", findLeastValueKey(map[int]int{
		1: 10,
		2: 40,
		3: 20,
		4: 5,
		5: 30,
	}))

	fmt.Printf("%v\n", findLeastValueKey(map[string]int{
		"a": 10,
		"b": 40,
		"c": 20,
		"d": 5,
		"e": 30,
	}))

	fmt.Printf("%v\n", findLeastValueKey(map[float64]int{
		1.0: 10,
		2.0: 40,
		3.0: 20,
		4.0: 50,
		5.0: 30,
	}))
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

type costs struct {
	day  int
	cost float64
}

func appendCostPerDay(costs []costs) []float64 {
	ans := []float64{}
	for i := 0; i < len(costs); i++ {
		for costs[i].day >= len(ans) {
			ans = append(ans, 0.0)
		}
		ans[costs[i].day] += costs[i].cost
	}
	return ans
}

func createMatrix(n int) [][]int {
	mat := make([][]int, 0)
	for i := 0; i < n; i++ {
		row := make([]int, 0)
		for j := 0; j < n; j++ {
			row = append(row, i*j)
		}
		mat = append(mat, row)
	}
	return mat
}

// ****************
// CHAPTER-9: Maps
// ****************

func countOccurances(chars string) map[rune]int {
	charMap := make(map[rune]int)
	for _, char := range chars {
		ch := rune(char)
		occr, ok := charMap[ch]
		if ok {
			charMap[ch] = occr + 1
		} else {
			charMap[ch] = 1
		}
	}
	return charMap
}

func deleteLeastOccurance(charMap map[rune]int) {
	leastOccr := 10
	for _, v := range charMap {
		if v < leastOccr {
			leastOccr = v
		}
	}
	for k, v := range charMap {
		if v == leastOccr {
			delete(charMap, k)
		}
	}
}

// ****************
// CHAPTER-10: Advance Functions
// ****************

// Higher Order Function
func messageFormatter(msgs []string, formatter func(string) string) []string {
	newMsgs := make([]string, len(msgs))
	for _, msg := range msgs {
		newMsgs = append(newMsgs, formatter(msg))
	}
	return newMsgs
}

// First Class Function
func formatString(str string) string {
	return fmt.Sprintf("Hello!\n%s\nRegards\n", str)
}

// Currying
func addNumbers(x, y int) int {
	return x + y
}

func multiplyNumbers(x, y int) int {
	return x * y
}

func doMath(calculate func(int, int) int) func(int) int {
	return func(x int) int {
		return calculate(x, x)
	}
}

type User struct {
	name  string
	admin bool
}

// defer - do something before the function ends - closing files, db connections
// multiple defer in same functions executes bottom-up order
func deleteUser(users map[string]User, name string) string {
	defer delete(users, name)
	if user, ok := users[name]; ok {
		if user.admin {
			return fmt.Sprintf("admin %s is deleted", name)
		} else {
			return fmt.Sprintf("user %s is deleted", name)
		}
	}
	return fmt.Sprintf("user %s not found", name)

}

// Clouser - A variable enclosed in a function
func stringBuilder() func(string) string {
	str := ""
	return func(word string) string {
		str += word + " "
		return str
	}
}

// Anonymous Function
func calculateCostOfMessages(msgs []string, cost func(string) float64) float64 {
	totalCost := 0.0
	for _, msg := range msgs {
		totalCost += cost(msg)
	}
	return totalCost
}

// ****************
// CHAPTER-11: Concurrency
// ****************

func sendEmail(message string) {
	go func() { // this won't get printed bcz program exits before go routine ends
		time.Sleep(time.Millisecond * 1000)
		fmt.Printf("Email received: %s\n", message)
	}()
	fmt.Printf("Email sent: %s\n", message)
}

// Channel
var emailSentChannel = make(chan bool)

func sendEmailUsingChannel(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: %s\n", message)
		emailSentChannel <- true
	}()
	fmt.Printf("Email sent: %s\n", message)
}

// Channel with range
func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func printFibonacci(n int) {
	ch := make(chan int)
	go func() {
		fibonacci(n, ch)
	}()
	fmt.Println("Fibonacci numbers")
	for num := range ch {
		fmt.Println(num)
	}
}

// Select - listening to multiple channels
func createNotification(messages, emails chan string) {
	go func() {
		for i := 0; i < 10; i++ {
			messages <- "message-" + strconv.Itoa(i)
			emails <- "email-" + strconv.Itoa(i)
			time.Sleep(time.Second)
		}
		close(messages)
		close(emails)
	}()
}
func showNotifications(messages, emails chan string) {
	for {
		select {
		case msg, ok := <-messages:
			if !ok {
				return
			}
			fmt.Printf("New message: %s\n", msg)
		case email, ok := <-emails:
			if !ok {
				return
			}
			fmt.Printf("New email: %s\n", email)
		}
	}
}

// ****************
// CHAPTER-12: Mutexes
// ****************

type userBalance struct {
	balance float64
	mux     *sync.RWMutex
}

func (ub *userBalance) credit(val float64) {
	ub.mux.Lock()
	defer ub.mux.Unlock()
	ub.balance += val
	time.Sleep(time.Second)
}

func (ub *userBalance) debit(val float64) {
	ub.mux.Lock()
	defer ub.mux.Unlock()
	ub.balance -= val
	time.Sleep(time.Second)
}

func (ub *userBalance) enquiry() {
	ub.mux.RLock()
	defer ub.mux.RUnlock()
	fmt.Println(ub.balance)
}

func doTransactions(ub userBalance) {
	var wb sync.WaitGroup
	for i := 0; i < 10; i++ {
		wb.Add(1)
		go func() {
			ub.credit(10)
			fmt.Println("Balance Credit Enquiry: ")
			ub.enquiry()
			wb.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		wb.Add(1)
		go func() {
			ub.debit(5)
			fmt.Println("Balance Debit Enquiry: ")
			ub.enquiry()
			wb.Done()
		}()
	}
	wb.Wait()
}

// ****************
// CHAPTER-13: Generics
// ****************

func splitAnySlice[T any](slice1 []T, n int) ([]T, []T, error) {
	var zeroVal []T
	if len(slice1) < n {
		return zeroVal, zeroVal, fmt.Errorf("Out of bounds")
	}
	if len(slice1) == 0 {
		return zeroVal, zeroVal, nil
	}
	return slice1[:n], slice1[n:], nil
}

type comparable interface {
	int | float64 | string
}

func findLeastValueKey[T comparable](m map[T]int) T {
	least := math.MaxInt
	var key T
	for k, v := range m {
		if least > v {
			least = v
			key = k
		}
	}
	return key
}

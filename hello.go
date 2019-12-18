/*

Credit to Mark McGranaghan and https://gobyexample.com/
Almost all of this is adapted/copied from there

Summary of Go:

** Pros
* Powerful and cool language features supporting high performance concurrency
	with the use of goroutines, channels, and select
* Easy, c-style syntax
* Static typing
* Compiles and runs really fast
* Garbage collected
* Used by big companies; Google, Docker...

** Cons
* Implicit interfaces (debatably a weakness or strength)
* No generics (to be added some day..?)
* Decent, but lacking library support
* Apparently the community is stubborn
* Fractured dependency management systems (go modules will one day be the standard?)

*/

package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	fmt.Printf("hello, world\n")
}

func doSomeLoopingAndConditionals() {
	// Variable declarations
	var i, j = 1, 2
	var k int = 3
	p := 4
	fmt.Println(j + k*p)

	// Go only has one looping keyword c:
	for i <= 3 {
		i = i + 1
	}

	for h := 7; h <= 9; h++ {
		if h%2 == 0 {
			continue
		} else {
			fmt.Println("yay")
		}
	}

	for {
		if result := i + j; result < 100 {
			fmt.Println("result is puny")
		} else {
			fmt.Println("result is large")
		}
		fmt.Println("loop test")
		break
	}

	switch p {
	case 1:
		fmt.Println("one")
	case 2, 3:
		fmt.Println("two")
	default: // optional
		fmt.Println("default")
	}

	// type switch to find type of interface
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func arraysAndSlices() {
	var myFirstArray [5]int
	myFirstArray[4] = 100
	fmt.Println(myFirstArray)

	secondArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(secondArray)

	var twoDArray [2][3]int // not true 2d, composed
	fmt.Println(twoDArray)

	/*
		Slices are more common than arrays in go
		They support additional functions like:
	*/
	s := make([]string, 3)
	s[0] = "a"

	// append (return a new slice with new element)
	s = append(s, "d")
	s = append(s, "d")
	s = append(s, "d")

	// copy
	c := make([]string, len(s))
	copy(c, s)

	// slice s operator (get slice in range)
	sliced := s[2:4]
	sliced = s[:5]
	fmt.Println(sliced)

	// multidimensional structure with variable column lengths
	twoDSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDSlice[i][j] = i + j
		}
	}
}

func maps() {
	// Maps; associative data type, AKA hashs or dicts
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	v1 := m["k1"]
	fmt.Println(v1)
	delete(m, "k2")

	// optionalSecondReturn contains info on whether key existed
	optionalSecondReturn, prs := m["k2"]
	fmt.Println(optionalSecondReturn, prs)
}

func ranges() {
	// used to iterate over several data structures
	// for example on arrays, slices, maps, and strings
	nums := []int{2, 3, 4}
	sum := 0

	// first return is index (or key for maps)
	// second return is value
	for _, num := range nums {
		sum += num
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}

///////////// Functions!

func addStuff(a int, b int) int {
	return a + b
}

func moreAdding(a, b, c int) int {
	return a + b + c
}

// Multi return
func multipleReturns() (int, int) {
	return 3, 7
}

func testMultipleReturns() {
	// Note, if you don't want all returns you
	// can use _ blank identifier
	a, b := multipleReturns()
	fmt.Println(a, b)
}

// Variadics
func variadicFunctionForSumming(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Print(nums, " ")
	fmt.Println(total)
}

func testVariadicFunction() {
	variadicFunctionForSumming(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	variadicFunctionForSumming(nums...)
}

// Closures (this shit is kewl)
func closureReturner() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func closureTester() {
	nextInt := closureReturner()
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	differentInts := closureReturner()
	fmt.Println(differentInts()) // 1
}

func recursiveFunction(n int) int {
	// factorial example
	if n == 0 {
		return 1
	}
	return n * recursiveFunction(n-1)
}

/////////// End Function stuff

///////////// Pointers

func takesVal(ival int) {
	ival = 0
}

func takesPtr(iptr *int) {
	*iptr = 0
}

func testPointers() {
	i := 1
	fmt.Println(i) // 1

	takesVal(i)
	fmt.Println(i) // 1

	takesPtr(&i)
	fmt.Println(i) // 0

	fmt.Println(&i) // Address in memory
}

///////////// Structs

type person struct {
	name string
	age  int
}

// NewPersonConstructor ; idiomatic to wrap struct init in constructor function
func NewPersonConstructor(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func testStructs() {
	bob := person{"Bob", 20}
	fmt.Println(bob) // {Bob 20}

	// Can have named args
	fmt.Println(person{name: "Chuck", age: 13}) // {Chuck 13}

	// Can have blank args if named
	fmt.Println(person{name: "Alice"}) // {Alice 0}

	fmt.Println(&person{name: "Ann", age: 40}) // &{Ann 40}
	fmt.Println(NewPersonConstructor("Jon"))   // &{Jon 42}

	fmt.Println(bob.age) // 20
	bob2 := &bob
	fmt.Println(bob2.age) // 20
	bob2.age = 99
	fmt.Println(bob.age) // 99
}

// struct with methods
type dog struct {
	name        string
	age, weight int
}

// receiver type of *dog
func (d *dog) yearOfBirth() int {
	return 2019 - d.age
}

// receiver type of *dog
func (d dog) healthFactor() int {
	return d.age * d.weight
}

// May want receiver type of value or ptr to avoid value copying or to allow modification of struct in function
func testMethodStruct() {
	doggy := dog{name: "Sam", age: 2, weight: 35}
	fmt.Println(doggy.healthFactor())
	fmt.Println(doggy.yearOfBirth())

	// Ptr to value conversions automatically handled by go
	doggyPtr := &doggy
	fmt.Println(doggyPtr.healthFactor())
	fmt.Println(doggyPtr.yearOfBirth())
}

/////////// Interface

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// Rectangle interface implementation
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Circle interface implementation
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func testInteraface() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// and... with a poof of smoke go figures out if your
	// structs implement the interface
	measure(r)
	measure(c)
}

//////// Errors

// by convention the last arg is of built in interface type "error"
// if a function can return an error
func functionWithDefaultError(arg int) (int, error) {
	if arg == 13 {
		return -1, errors.New("unlucky number detected")
	}
	return arg + 1, nil // nil means no error
}

// Can define custom errors if they implement the Error() method
type customError struct {
	arg  int
	prob string
}

func (e *customError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func functionWithCustomError(arg int) (int, error) {
	if arg == 13 {
		return -1, &customError{arg, "just can't do it bro"}
	}
	return arg + 2, nil
}

func testErrors() {
	for _, i := range []int{1, 13} {
		if r, e := functionWithDefaultError(i); e != nil {
			fmt.Println("default err func failed:", e)
		} else {
			fmt.Println("default err func win:", r)
		}
	}

	for _, i := range []int{1, 13} {
		if r, e := functionWithCustomError(i); e != nil {
			fmt.Println("custom err func failed:", e)
		} else {
			fmt.Println("custom err func win:", r)
		}
	}

	// This how to cast an error to use its data
	_, e := functionWithCustomError(13)
	if castedError, castAssertionPassed := e.(*customError); castAssertionPassed {
		fmt.Println(castedError.arg)
		fmt.Println(castedError.prob)
	}
}

/////////// GoRoutines
// A lightweight thread of execution
// It is truly concurrent and can utilize separate cores on your machine (unline in say, python)

func somethingToRun(name string, loops int) {
	for i := 0; i < loops; i++ {
		fmt.Println(name, ":", i)
	}
}

func testGoRoutines() {
	somethingToRun("sync", 3)
	go somethingToRun("async", 5)

	// anon function in a goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("HELLO")
}

/////////// Channels
// pipes that pass information between concurrent goroutines
// Note, timers and tickers (wait for timer, or do something repeating on a ticker) can be implementing using channels (and select for tickers)... https://gobyexample.com/tickers
// Note, worker pools can be easily implemented with channels... https://gobyexample.com/worker-pools
func testChannels() {
	//// Normal channel
	basicChannel := make(chan string)

	go func() { basicChannel <- "ping" }() // put message in channel

	// Note the go routine is by default locked until the receiver accepts message
	msg := <-basicChannel // receive message from channel
	fmt.Println(msg)      // prints: ping

	//// Buffered channel
	// This channel accepts 2 values before locking its thread
	bufferedChannel := make(chan string, 2)
	bufferedChannel <- "buffered"
	bufferedChannel <- "channel"
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
}

//// Sync threads with channels
// Note, syncing multiple goroutines may be better done with a WaitGroup
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func testSyncWithWorker() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}

//// Directional type safety for channels
func ping(insertOnlyChannel chan<- string, msg string) {
	insertOnlyChannel <- msg
}

func pong(insertOnlyChannel chan<- string, popOnlyChannel <-chan string) {
	msg := <-popOnlyChannel
	insertOnlyChannel <- msg
}

func testChannelDirections() {
	channelOne := make(chan string, 1) // Note, the 1 here makes this channel not block when only 1 value is in it!
	channelTwo := make(chan string, 1)
	ping(channelOne, "my sweet message")
	pong(channelTwo, channelOne)
	fmt.Println(<-channelTwo)
}

//// Select
// lets you wait on multiple channels
func testSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Simultaneously wait for both channels and print each when ready
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-time.After(5 * time.Second): // Timeouts are easy with select!
			fmt.Println("TIMEOUT")
		}
	}
}

func testNonBlockingChannelsWithSelect() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	// Non blocking read
	select {
	case msg := <-channel1: // If tehre is a message ready, take it, otherwise go default
		fmt.Println("first msg", msg)
	case msg2 := <-channel2: // You can do multiple non blocking reads and writes
		fmt.Println("second msg", msg2)
	default:
		fmt.Println("nothing here")
	}

	// Non blocking write
	select {
	case channel2 <- "hi": // Send message if receiver is ready, default otherwise
		fmt.Println("sent")
	default:
		fmt.Println("no one to receive")
	}

}

func testClosingChannels() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs // more is true unless channel is closed
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) // close the channel! Note, a closed channel can still have its values received
	fmt.Println("sent all jobs")

	<-done
	fmt.Println("Done processing jobs")
}

func showRangeWithChannels() {
	queue := make(chan string, 3)
	queue <- "1"
	queue <- "2"
	queue <- "3"
	close(queue)

	// Iterate over values in a channel
	for elem := range queue {
		fmt.Println(elem)
	}
}

////////// Panic
// Quickly exit a program if an error is received that you don't know how or want to handle
func showPanic() {
	panic("a problem")

	// Common error non handling pattern
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

///////// Defer
// do something at the end of the enclosing function (kind of like 'finally' in other languages)

func testFinally() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f) // Execute when this enclosing function ends
	writeFile(f)

	// Note defer wont be called if a panic happens before end of function
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	// Note, you should still check for errors when closing files even if its in a deferred function
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

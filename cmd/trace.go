package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

type interface1 interface {
	M1(int)
	M11()
}

type interface2 interface {
	M1(int)
	M21()
}

type Empty interface {
}

type interface3 interface {
	interface1
	interface2
}

func Trace() func() {

	pc, fileName, lineNumber, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("not found caller.")
	}

	funcc := runtime.FuncForPC(pc)
	funcName := funcc.Name()

	fmt.Println("enter func : ", funcName, fileName, lineNumber)
	return func() {
		fmt.Println("leave func : ", funcName)
	}

}

func normalDefer() {
	fmt.Println("normal")
}

func foo() {
	defer normalDefer()
	fmt.Println("excuate func foo beforce defer.")
	defer Trace()()
	fmt.Println("excuate func foo after defer.")
	bar()

}

func bar() {
	go curGoroutineID()

	curGoroutineID()
	fmt.Println("excuate func bar beforce defer.")
	defer Trace()()
	fmt.Println("excuate func bar after defer.")

}

// trace2/trace.go
var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}

	fmt.Println("Goruntine is ", n)
	return n
}

func interfaceT(empty interface{}) {
	fmt.Println(empty)

	var a Empty
	xx, ok := a.(int)
	fmt.Println(xx, ok)

	switch empty.(type) {
	case int:
		fmt.Println("parameter is int. ")

	case string:
		fmt.Println("parameter is string.")

	case []int:
		fmt.Println("parameter is slice. ")

	case [3]int:
		fmt.Println("parameter is array 3 int. ")

	default:
		fmt.Println("unkonwn type")
	}

	return
}

func assertInterface(empty interface{}) {
	v, ok := empty.(int)
	fmt.Printf("type is %T, value is %v, bool : %v\n", v, v, ok)

	var b any
	b = 10
	b = "123"
	fmt.Println(b)

}

func main() {
	defer Trace()()
	foo()

	var b [3]byte
	b[1] = 10
	b[2] = 20

	var d [4]int = [4]int{3, 4, 4, 5}
	fmt.Println(d)

	c := [3]byte{'a', 'b', 'c'}

	e := []byte{'c', 87, 96}
	fmt.Println(e)

	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("learn trace. ")

	interfaceT("abc")

	interfaceT(123)

	interfaceT([3]int{})

	interfaceT([4]int{})

	assertInterface(3)

	assertInterface("abc")

	var tmp interface{} = 123456
	fmt.Println(tmp)
}

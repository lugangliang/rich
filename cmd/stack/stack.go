package main

import (
	"errors"
	"fmt"
	"net/http"
)

func funcA() {
	funcB()

}

func funcB() {
	funcA()
}

type myinterface interface {
	M1(int)
}

type T struct {
}

func (t T) M1(a int) {

}

func (t T) Error() string {
	return "in T string"
}

type myError struct {
	error
}

func returnError() error {
	var p *myError = nil
	fmt.Println(p == nil)
	return p
}

func greeting(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w, r)
	return
}

func (t T) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, r)
	w.Write([]byte{97, 98, 101, 102})
}

func main() {

	//funcA()

	a := [3]int{1, 2, 3}
	b := &a
	fmt.Printf("a is : %v, type is : %T, address is %d \n", a, a, a)
	fmt.Printf("b is : %v, type is : %T, address is %p\n", b, b, b)
	fmt.Println(a[2], b[2], a[1], b[:])

	var c interface{}
	var t T
	c = t
	_, ok := c.(myinterface)
	if !ok {
		panic("interface not ")
	}

	var err error
	err = errors.New("hehe")
	fmt.Printf("type is %T, value is %v \n", err, err)

	err = t

	fmt.Printf("type is %T, value is %v \n", err, err)

	p := returnError()
	fmt.Println(p == nil)

	err = http.ListenAndServe(":8000", T{})
	fmt.Println(err)

	return
}

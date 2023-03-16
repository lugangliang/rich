package main

import (
	"errors"
	"fmt"
)

type QuackAble interface {
	Quack()
}

type Dog struct {
}

func (dog Dog) Quack() {
	fmt.Println("dog quack")
}

type Bird struct {
}

func (bird Bird) Quack() {
	fmt.Println("bird quack")
}

func interfaceTest() {
	var b any
	var e error

	fmt.Println("before :")

	fmt.Println(b, e)
	fmt.Println(b == nil)
	fmt.Println(e == nil)
	fmt.Println(b == e)

	fmt.Println("After : ")
	b = 13
	e = errors.New("")
	fmt.Println(b, e)
	fmt.Println(b == nil)
	fmt.Println(e == nil)
	fmt.Println(b == e)

}

func test2() {
	var n int = 16
	var ei interface{} = &n
	fmt.Println(n, ei)
	n = 17
	fmt.Println(n, ei)
}

func main() {

	animals := []QuackAble{&Dog{}, new(Bird)}

	for _, animal := range animals {
		animal.Quack()
	}

	//interfaceTest()
	test2()

	fmt.Println("go core : interface")
}

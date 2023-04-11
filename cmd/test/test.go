package main

import (
	"fmt"
	"time"
)

func cal() {
	sum := 0

	for i := 1; i <= 100000000000; i++ {
		sum += i
	}
	fmt.Println("1 add to 100 is ", sum)
}

func main() {

	for i := 1; i < 30; i++ {
		go cal()
	}

	time.Sleep(time.Minute)

	fmt.Println("Learn data .")
}

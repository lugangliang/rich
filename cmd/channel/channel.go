package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func spwan() <-chan error {
	c := make(chan error, 10)

	go func() {
		time.Sleep(3 * time.Second)
		c <- errors.New("Execute in spawn gorutine.")
	}()

	return c
}

func contextTest() {
	ctx, cancel := context.WithCancel(context.Background())

	chanel := make(chan string, 10)

	go func() {
		fmt.Println("go routine.")
		for {
			select {
			case <-ctx.Done():
				chanel <- "go routine exit."
				fmt.Println("nonono")
				return
			default:
				time.Sleep(300 * time.Millisecond)
				fmt.Println("default")
			}

		}
	}()

	time.Sleep(3 * time.Second)
	cancel()

	fmt.Println(<-chanel)

}

func deadloop() {
	for {

	}
}

func consume(cs <-chan int) {
	for c := range cs {
		time.Sleep(time.Second)
		fmt.Println("recieved ", c)
	}
}

func consume1(cs <-chan int) {

	time.Sleep(time.Second)
	fmt.Println("recieved ", <-cs)
}

func consume2(cs <-chan int) {

	for {
		time.Sleep(time.Second)
		a, ok := <-cs
		if ok {
			fmt.Println("recieved ", a)
		} else {
			return
		}
	}

}

func consume3(cs <-chan int) {
	for {
		select {
		case a, ok := <-cs:
			if !ok {
				return
			}
			time.Sleep(300 * time.Millisecond)
			fmt.Println("recieved ", a)
		}
	}

}

func producer(c chan<- int) {
	for i := 0; i < 30; i++ {
		fmt.Println("producer send ", i)
		c <- i
		time.Sleep(100 * time.Millisecond)
	}

	defer close(c)

	return
}

func syncTest() {
	chanel := make(chan int, 100)

	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println(&wg)

	go func() {
		producer(chanel)
		wg.Done()
	}()

	go func() {
		// consume(chanel)
		consume3(chanel)
		wg.Done()
	}()

	wg.Wait()
}

func worker(i int) {
	fmt.Printf("worker %d is working...\n", i)
	time.Sleep(time.Second)
	fmt.Printf("worker %d is done...\n", i)

	return
}

func spawnGroup() {
	var wg sync.WaitGroup

	var groupChanel = make(chan any)
	var exitChannel = make(chan any)

	for num := 0; num < 20; num++ {
		wg.Add(1)
		go func(i int) {
			<-groupChanel
			fmt.Printf("worker %d start to work .\n", i)
			worker(i)
			wg.Done()
		}(num)
	}

	go func() {
		wg.Wait()
		exitChannel <- 0
	}()

	close(groupChanel)
	<-exitChannel
	fmt.Println("all workers has done.")
	time.Sleep(3 * time.Second)

}

func groupTest() {
	spawnGroup()

}

type Counter struct {
	sync.Mutex
	sum int
}

func Increase(i int) int {
	sum.Lock()
	defer sum.Unlock()
	sum.sum++
	fmt.Printf("goroutine-%d is %d.\n", i, sum.sum)
	return sum.sum
}

var sum Counter

func sum1() {

	for i := 1; i < 10; i++ {
		go func(i int) {
			Increase(i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Final sum is ", sum)
}

func main() {

	runtime.GOMAXPROCS(10)

	fmt.Println("Learn go channel")

	//c := spwan()
	//fmt.Printf("%s\n", <-c)

	//contextTest()

	//go deadloop()

	// syncTest()

	// groupTest()
	sum1()
	time.Sleep(100 * time.Second)
	return
}

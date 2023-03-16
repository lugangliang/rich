package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

type Pool struct {
	capacity int
	tasks    chan Task
	actives  chan any
	wg       sync.WaitGroup
}

func (pool *Pool) Scheduler(f Task) {
	fmt.Println("submit a task.")
	pool.tasks <- f
}

func NewPool(capacity int) *Pool {

	var p = &Pool{
		capacity: capacity,
		tasks:    make(chan Task),
		actives:  make(chan any, capacity),
		wg:       sync.WaitGroup{},
	}

	return p
}

func (pool *Pool) newWorker() {
	fmt.Println("new goroutine")
	pool.wg.Add(1)
	go func() {
		for {
			select {
			case f, ok := <-pool.tasks:
				if !ok {
					fmt.Println("no task, contineue")
					continue
				}
				f()
			default:
				time.Sleep(time.Second)
			}

		}
		fmt.Println("go routine exit.")
		pool.wg.Done()
	}()

}

func (pool *Pool) Run() {

	for {
		select {
		case pool.actives <- struct{}{}:
			fmt.Println("new a worker.")
			pool.newWorker()
		default:
			time.Sleep(time.Second)
		}
	}

}

func main() {

	p := NewPool(5)

	go func() {
		for i := 0; i < 40; i++ {
			time.Sleep(time.Second)
			p.Scheduler(func() {
				fmt.Println("i'm working.")
			})
		}
	}()

	p.Run()

	p.wg.Wait()

	fmt.Println("Learn pool.")
}

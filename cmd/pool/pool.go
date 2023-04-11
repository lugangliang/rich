package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Task func(num int)

const (
	DefaultCapacity = 100
	MaxCapacity     = 200
)

type Pool struct {
	Prealloc bool
	block    bool
	capacity int
	tasks    chan Task
	actives  chan any
	wg       sync.WaitGroup
	quit     chan any
}

func (pool *Pool) Scheduler(f Task) {
	pool.tasks <- f
}

func NewPool(capacity int, opts ...Option) *Pool {

	if capacity <= 0 {
		capacity = DefaultCapacity
	} else if capacity >= MaxCapacity {
		capacity = MaxCapacity
	}

	var p = &Pool{
		capacity: capacity,
		tasks:    make(chan Task),
		actives:  make(chan any, capacity),
		wg:       sync.WaitGroup{},
		quit:     make(chan any),
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

type Option func(pool *Pool)

func (pool *Pool) newWorker(idx int) {
	fmt.Println("Create a new goroutine.")
	pool.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker[%03d]: recover panic and exit , %s\n", idx, err)
				<-pool.actives
			}
		}()
	loop:
		for {
			select {
			case <-pool.quit:
				fmt.Printf("main goroutine exit, goroutine %d exit\n", idx)
				return
			case f, ok := <-pool.tasks:
				if !ok {
					fmt.Println("no task, contineue")
					break loop
				}
				fmt.Printf("worker[%03d] execute a task \n", idx)
				f(idx)
			default:
				time.Sleep(time.Second)
			}

		}
		fmt.Printf("worker[%03d]go routine exit.", idx)
		pool.wg.Done()
	}()

}

func (pool *Pool) Run() {

	defer func() {
		pool.wg.Wait()
	}()

	idx := 0

	for {
		select {
		case _, ok := <-pool.quit:
			if !ok {
				fmt.Println("exit all goroutine.")
				return
			}

		case pool.actives <- struct{}{}:
			fmt.Printf("worker[%03d] receive a task.\n", idx)
			pool.newWorker(idx)
			idx++
		default:
			time.Sleep(time.Second)
		}
	}
}

func (pool *Pool) Free() {

	time.Sleep(10 * time.Second)
	fmt.Println("ready to broadcast quit signal.")
	close(pool.quit)

}

func WithPrealloc(prealloc bool) Option {
	return func(pool *Pool) {
		pool.Prealloc = prealloc
	}
}

func WithBlock(block bool) Option {
	return func(pool *Pool) {
		pool.block = block
	}
}

func main() {

	http.Transport{
		Proxy:                  nil,
		DialContext:            nil,
		Dial:                   nil,
		DialTLSContext:         nil,
		DialTLS:                nil,
		TLSClientConfig:        nil,
		TLSHandshakeTimeout:    0,
		DisableKeepAlives:      false,
		DisableCompression:     false,
		MaxIdleConns:           0,
		MaxIdleConnsPerHost:    0,
		MaxConnsPerHost:        0,
		IdleConnTimeout:        0,
		ResponseHeaderTimeout:  0,
		ExpectContinueTimeout:  0,
		TLSNextProto:           nil,
		ProxyConnectHeader:     nil,
		GetProxyConnectHeader:  nil,
		MaxResponseHeaderBytes: 0,
		WriteBufferSize:        0,
		ReadBufferSize:         0,
		ForceAttemptHTTP2:      false,
	}
	p := NewPool(5, WithBlock(true), WithPrealloc(false))

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			p.Scheduler(func(num int) {
				fmt.Println("  my task is ", num)
				if num%3 == 1 {
					panic("test")
				}
			})
		}
	}()

	go p.Run()

	time.Sleep(time.Second)

	p.Free()

	fmt.Println("Learn pool.")
}

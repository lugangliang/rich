package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"sync"
	"time"
)

var logger *zap.Logger

var a int

var mutex sync.Mutex

type myFunc func(int) (int, string)

func init() {
	logger, _ = zap.NewProduction()
	logger.Info("ss")
	mutex.Lock()
	mutex.Unlock()
	mutexCopy := &mutex
	mutexCopy.Lock()
	mutexCopy.Unlock()
}

func funcTest(myFunc2 myFunc) {

	fmt.Println("funcTest")
	myFunc2(10)

}

func deferTest() {

	for i := 0; i < 4; i++ {
		defer fmt.Println("value is ", i)
	}

	return
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	logger.Info("hello, go module", zap.ByteString("uri", ctx.RequestURI()))
}

func main() {

	logrus.Println("hello, logrus")
	logrus.Println(uuid.NewString())

	sSlice := []int{2, 3, 4, 5, 6}

	deferTest()

	fmt.Errorf("abcdefg")

	funcTest(func(int) (int, string) { return 1, "abc" })

	for index, value := range sSlice[:] {
		//indexCopy := index
		//valueCopy := value
		go func(index int, value int) {
			time.Sleep(time.Second)
			//fmt.Printf("update : %d %x %d %x\n", index, &index, value, &value)
			fmt.Printf("slice address : %p\n", sSlice)
			sSlice[2] = 100
			//fmt.Printf("origin : %d %x %d %x\n", indexCopy, &indexCopy, valueCopy, &valueCopy)
		}(index, value)
	}
	time.AfterFunc(5*time.Second, func() { fmt.Println("time fired") })

	fmt.Printf("original slice address : %p\n", sSlice)
	time.Sleep(3 * time.Second)
	fmt.Println(sSlice)
	fasthttp.ListenAndServe(":8081", fastHTTPHandler)
	fmt.Println("hello world")

	fmt.Println(a)

	return
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	_ "rich/internal/store"
	"rich/server"
	"rich/store/factory"
	"syscall"
)

func main() {

	s, err := factory.New("mem")
	if err != nil {
		fmt.Print(err)
		return
	}

	bServer, _ := server.NewBookStoreServer(":8080", s)
	errChan, startErr := bServer.ListenAndServe()
	if startErr != nil {
		fmt.Println(errChan)
	}

	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-c:
		fmt.Println("Recieve signal")
		return
	case <-errChan:
		fmt.Println("server start failed, ", startErr)
	}

	fmt.Println("program exit.")

	return
}

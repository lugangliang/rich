package main

import (
	"fmt"
	_ "rich/internal/store"
	"rich/store/factory"
)

func main() {

	s, err := factory.New("mem")
	if err != nil {
		fmt.Print(err)
		return
	}

	_, err = s.Get("haha")
	if err != nil {
		fmt.Print(err)
	}

	return
}

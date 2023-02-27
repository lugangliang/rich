package factory

import (
	"errors"
	"fmt"
	"rich/store"
	"sync"
)

var (
	lock         sync.Mutex
	storeFactory = make(map[string]store.Store)
)

func init() {
	fmt.Println("factory init ...")
}

func Register(name string, s store.Store) {
	lock.Lock()
	defer lock.Unlock()
	if _, dup := storeFactory[name]; dup {
		fmt.Printf("dup name : %s\n", name)
		panic("dup")
	}

	storeFactory[name] = s
}

func New(name string) (s store.Store, err error) {

	if _, exist := storeFactory[name]; exist {
		return storeFactory[name], nil
	}

	return nil, errors.New(fmt.Sprintf("Driver type %s is not found. ", name))
}

package store

import (
	"errors"
	"fmt"
	mystore "rich/store"
	storefactory "rich/store/factory"
)

type MemStore struct {
	books map[string]*mystore.Book
}

func (ms *MemStore) Create(name string, book *mystore.Book) error {
	ms.books[name] = book

	return nil
}

func (ms *MemStore) Delete(name string) error {

	delete(ms.books, name)

	return nil

}

func (ms *MemStore) GetAll() ([]*mystore.Book, error) {

	var allBooks []*mystore.Book

	for _, Book := range ms.books {
		allBooks = append(allBooks, Book)
	}

	return allBooks, nil
}

func (ms *MemStore) Get(name string) (book *mystore.Book, err error) {
	if _, exist := ms.books[name]; !exist {
		return nil, errors.New("Not found")
	}
	return ms.books[name], nil
}

func (ms *MemStore) Update(name string, newBook *mystore.Book) error {
	ms.books[name] = newBook

	return nil
}

func init() {
	fmt.Println("Register memory store to factory.")
	storefactory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

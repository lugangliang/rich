package main

import (
	"fmt"
)

type SingleList struct {
	data int
	next *SingleList
}

func newNode(data int) *SingleList {
	tmp := &SingleList{
		data: data,
		next: nil,
	}

	return tmp
}

func (list *SingleList) insertNodeToHead(node *SingleList) {
	if list.next == nil {
		fmt.Printf("%p\n", &list)
		list.next = node
		return
	}

	node.next = list.next
	list.next = node
}

func (list *SingleList) insertNodeToEnd(node *SingleList) {
	if list.next == nil {
		list.next = node
		return
	}

	tmp := list
	for tmp != nil {
		if tmp.next == nil {
			tmp.next = node
			return
		}
		tmp = tmp.next
	}

}

func (list *SingleList) reserveSingleList() {

	newHead := newNode(0)

	if list.next == nil {
		return
	}

	if list.next.next == nil {
		return
	}

	tmp := list.next
	for tmp != nil {
		newHead.insertNodeToHead(tmp)
		tmp = tmp.next
	}

	list.next = newHead.next

}

func newList() *SingleList {
	head := newNode(0)
	for i := 0; i < 20; i++ {
		tmp := newNode(i)
		(head).insertNodeToHead(tmp)
	}
	return head
}

func (list *SingleList) PrintList() {
	if list.next == nil {
		fmt.Println("empty list.")
	}

	tmp := list
	for tmp != nil {
		fmt.Printf(" %d,", tmp.data)
		tmp = tmp.next
	}
	fmt.Println()
}

func main() {
	head := newList()
	head.PrintList()
	head.reserveSingleList()
	head.PrintList()
	fmt.Println("Learn single list.")
}

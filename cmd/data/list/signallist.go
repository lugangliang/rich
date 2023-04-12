package main

import (
	"fmt"
	"time"
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
		list.next = node
		node.next = nil
		return
	}

	node.next = list.next
	list.next = node
}

func (list *SingleList) insertNodeToEnd(node *SingleList) {
	if list.next == nil {
		list.next = node
		node.next = nil
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

func (list *SingleList) reserveSingleListMethod1() {

	newHead := newNode(0)

	// empty list or only one node.
	if list.next == nil || list.next.next == nil {
		return
	}

	for tmp := list.next; tmp != nil; {
		// record tmp's next pointer
		next := tmp.next
		newHead.insertNodeToHead(tmp)
		tmp = next
	}

	list.next = newHead.next
}

func (list *SingleList) reserveSingleListMethod2() {
	// empty list or only one node.
	if list.next == nil || list.next.next == nil {
		return
	}

	secondElem := list.next

	// 依次将下一个元素的下一个元素插到头上去。
	for tmp := list.next.next; tmp != nil; {
		next := tmp.next
		secondElem.next = next
		list.insertNodeToHead(tmp)
		tmp = next
	}
}

func newList() *SingleList {
	head := newNode(0)
	for i := 2; i < 30; i++ {
		tmp := newNode(i)
		(head).insertNodeToHead(tmp)
	}
	return head
}

func (list *SingleList) PrintList() {
	if list.next == nil {
		fmt.Println("empty list.")
	}

	// ignore first element.
	tmp := list.next
	for tmp != nil {
		fmt.Printf(" %d,", tmp.data)
		tmp = tmp.next
	}
	fmt.Println()
}

func main() {
	head := newList()
	head.PrintList()

	time.Sleep(2 * time.Second)
	head.reserveSingleListMethod1()
	head.PrintList()

	time.Sleep(2 * time.Second)
	head.reserveSingleListMethod2()
	head.PrintList()
	fmt.Println("Learn single list.")
}

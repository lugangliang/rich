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

var record = map[int]int{
	1: 1,
	2: 2,
}

func grow1(n int) int {

	if _, ok := record[n]; ok {
		fmt.Printf("num %d is processed, value is %d .\n", n, record[n])
		return record[n]
	}

	fmt.Printf("num %d is start processin .\n", n)

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	resn1 := grow1(n - 1)
	resn2 := grow1(n - 2)
	res := resn1 + resn2

	record[n] = res
	fmt.Printf("num %d is processed, value is %d from lower .\n", n, record[n])

	return res
}

func grow2(n int) int {
	fmt.Printf("num %d is start processin .\n", n)

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	resn1 := grow2(n - 1)
	resn2 := grow2(n - 2)
	res := resn1 + resn2
	fmt.Printf("num %d is processed, value is %d from lower .\n", n, res)

	return res
}

func num(n int) int {
	if n == 1 {
		return 1
	}

	return num(n-1) + 1
}

func test() {
	head := newList()
	head.PrintList()

	time.Sleep(2 * time.Second)
	head.reserveSingleListMethod1()
	head.PrintList()

	time.Sleep(2 * time.Second)
	head.reserveSingleListMethod2()
	head.PrintList()
}

func main() {
	//fmt.Println(grow2(10))
	//fmt.Println(num(6))

	fmt.Println("Learn single list.")
}

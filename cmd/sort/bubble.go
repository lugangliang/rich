package main

import (
	"fmt"
)

var arr = [10]int{40, 5, 2, 10, 7, 30, 6, 3, 21, 1}

func init() {
	fmt.Println("orginal : ")
	fmt.Println("   ", arr)

	sortMethod["bubble"] = bubble
	sortMethod["insert"] = insert

}

func insert() {

	length := len(arr)

	var value int

	// 从未排序区间取第一个元素 插入到已排序区间
	for i := 1; i < length; i++ {
		value = arr[i]
		j := i - 1
		// 依次与已排序区间的每个元素进行比较
		for ; j >= 0; j-- {
			// 找到应该插入的位置,i 是待排序的数据，j是从大到0，依次与i的数值进行比较
			if arr[j] > value {
				// i 应该在j的位置, j 到 i-1 之间的元素后移
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = value
	}

	fmt.Println("after insert")
	fmt.Println("   ", arr)
	fmt.Println("insert sort ")
}

func bubble() {
	length := len(arr)

	var ok bool

	for i := 0; i < length; i++ {
		ok = false
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				ok = true
			}
		}

		if ok == false {
			fmt.Println("heree ")
			break
		}
	}

	fmt.Println("after sort:")
	fmt.Println("   ", arr)

}

var sortMethod = make(map[string]func())

func main() {

	sortMethod["insert"]()

	fmt.Println("Learn bubble sort.")
}

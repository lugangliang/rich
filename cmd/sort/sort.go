package main

import (
	"fmt"
	"math/rand"
	"time"
)

var arr []int

func genRandNum() {
	for i := 0; i < 21; i++ {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Microsecond)
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr, len(arr), cap(arr))
}

// 奇数个元素
// 偶数个元素
// 有相等元素
// 空元素
// 1个元素
// 2个元素

func partion(arr []int, left int, right int) int {

	value := arr[right]

	i := left

	for j := left; j < right; j++ {
		if arr[j] < value {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
		time.Sleep(200 * time.Millisecond)
		fmt.Println(i, j)
	}

	arr[i], arr[right] = arr[right], arr[i]

	return i
}

func quickSort1(arr []int) {

	quickSortSub1(arr, 0, len(arr)-1)

}

func quickSortSub1(arr []int, left int, right int) {

	if left >= right {
		return
	}

	pivot := partion(arr, left, right)

	quickSortSub1(arr, left, pivot-1)
	quickSortSub1(arr, pivot+1, right)

}

func main() {

	genRandNum()

	quickSort1(arr)

	fmt.Println(arr)
	fmt.Println("Learn sort.")
}

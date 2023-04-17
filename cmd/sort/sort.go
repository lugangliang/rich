package main

import (
	"fmt"
	"math/rand"
	"time"
)

var quickArr []int

func genRandNum() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(33)
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Microsecond)
		quickArr = append(quickArr, rand.Intn(100))
	}
	fmt.Println(quickArr, len(quickArr), cap(quickArr))
}

// 奇数个元素
// 偶数个元素
// 有相等元素
// 空元素
// 1个元素
// 2个元素

func partion(quickArr []int, left int, right int) int {

	value := quickArr[right]

	i := left

	for j := left; j < right; j++ {
		if quickArr[j] < value {
			quickArr[j], quickArr[i] = quickArr[i], quickArr[j]
			i++
		}
	}

	quickArr[i], quickArr[right] = quickArr[right], quickArr[i]

	return i
}

func quickSort1(arr []int) {

	quickSortSub1(arr, 0, len(arr)-1)

}

func quickSortSub1(Arr []int, left int, right int) {

	if left >= right {
		return
	}

	pivot := partion(Arr, left, right)

	quickSortSub1(Arr, left, pivot-1)
	quickSortSub1(Arr, pivot+1, right)

}

func check() {
	fmt.Println(quickArr, len(quickArr), cap(quickArr))

	for i := 0; i < len(quickArr)-1; i++ {
		if quickArr[i] > quickArr[i+1] {
			fmt.Println("error")
			return
		}
	}

	fmt.Println("quick sort works")
}

func init() {
	genRandNum()
}

func main() {

	quickSort1(quickArr)

	check()
	fmt.Println("Learn sort.")
}

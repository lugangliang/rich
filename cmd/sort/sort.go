package main

import (
	"fmt"
	"math/rand"
	"time"
)

var quickArr []int

func genRandNum() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(200)
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Microsecond)
		quickArr = append(quickArr, rand.Intn(1000))
	}
	fmt.Println(quickArr, len(quickArr), cap(quickArr))
}

// 奇数个元素
// 偶数个元素
// 有相等元素
// 空元素
// 1个元素
// 2个元素

func partionTail(quickArr []int, left int, right int) int {

	value := quickArr[right]

	i := left

	for j := left; j < right; {
		fmt.Printf("	before compare arr=%v, i=%d,j=%d \n", quickArr, i, j)
		if quickArr[j] < value {
			quickArr[j], quickArr[i] = quickArr[i], quickArr[j]
			i++
		}
		j++
		fmt.Printf("	after compare arr=%v, i=%d, j=%d\n", quickArr, i, j)

	}

	quickArr[i], quickArr[right] = quickArr[right], quickArr[i]

	fmt.Printf("final result arr=%v, i=%d, left=%d, right=%d\n", quickArr, i, left, right)

	return i
}

func partionHead(arr []int, left int, right int) int {

	value := arr[left]

	for left < right {

		// 从右到左，找到一个比value小的元素，交换到左边，右边出来一个坑
		for left < right {
			if arr[right] < value {
				arr[left] = arr[right]
				break
			} else {
				right--
			}
		}

		// 从左到右，找到一个比value小的元素，交换到左边
		for left < right {
			if arr[left] > value {
				arr[right] = arr[left]
				break
			} else {
				left++
			}
		}
	}

	// 两个指针碰到一起了，
	fmt.Println("left == right is ", left == right)
	quickArr[left] = value

	return left
}

func quickSortTail(arr []int) {

	quickSortSubTail(arr, 0, len(arr)-1)

}

func quickSortHead(arr []int) {

	quickSortSubHead(arr, 0, len(arr)-1)

}

func quickSortSubTail(Arr []int, left int, right int) {
	fmt.Printf("compare arr %v , left=%d, right=%d\n", Arr, left, right)
	if left >= right {
		return
	}

	pivot := partionTail(Arr, left, right)

	quickSortSubTail(Arr, left, pivot-1)
	quickSortSubTail(Arr, pivot+1, right)

}
func quickSortSubHead(Arr []int, left int, right int) {
	fmt.Printf("compare arr %v , left=%d, right=%d\n", Arr, left, right)
	if left >= right {
		return
	}

	pivot := partionHead(Arr, left, right)

	quickSortSubHead(Arr, left, pivot-1)
	quickSortSubHead(Arr, pivot+1, right)

}

func check() {
	fmt.Println(quickArr, len(quickArr), cap(quickArr))

	for i := 0; i < len(quickArr)-1; i++ {
		if quickArr[i] > quickArr[i+1] {
			fmt.Println("error")
			panic("error")
			return
		}
	}

	fmt.Println("quick sort works")
}

func init() {
	genRandNum()
}

func main() {

	for i := 0; i < 100; i++ {
		quickArr = []int{}
		genRandNum()
		quickSortTail(quickArr)
		//quickSortHead(quickArr)

		check()
		//time.Sleep(5 * time.Second)
	}

	fmt.Println("Learn sort.")
}

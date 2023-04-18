package main

import (
	"fmt"
	"time"
)

// 找不到元素
// 奇数元素
// 偶数元素
// 空元素
// 1个元素
// 2个元素
// 3个元素
// 4个元素

var bArr [][]int

func init() {
	arr7 := [7]int{3, 4, 8, 10, 20, 40, 60}

	arr8 := [8]int{3, 4, 8, 10, 20, 40, 60, 77}

	arr2 := [2]int{3, 4}

	arr3 := [3]int{3, 4, 5}

	arr4 := [4]int{4, 10, 22, 33}

	bArr = append(bArr, arr2[:])
	bArr = append(bArr, arr3[:])
	bArr = append(bArr, arr3[:])
	bArr = append(bArr, arr7[:])
	bArr = append(bArr, arr8[:])

	fmt.Println(arr2, arr3, arr4, arr7, arr8)
	fmt.Println(bArr)
}
func bsearch(arr []int, value int) (bool, int) {
	fmt.Println(arr)
	lenght := len(arr)

	var low = 0
	var high = lenght - 1
	var mid = low + (high-low)>>1

	for {

		if low > high {
			return false, 0
		} else if low == high {
			fmt.Println("low high is equal.")
		}

		time.Sleep(time.Second)

		fmt.Println(low, mid, high)

		if value == arr[mid] {
			return true, mid
		} else if value > arr[mid] {
			low = mid + 1
			mid = low + (high-low)>>1
			continue
		} else {
			high = mid - 1
			mid = low + (high-low)>>1
			continue
		}
	}

}

func main() {

	for _, arr := range bArr {
		fmt.Println(bsearch(arr, 77))
	}

	fmt.Println("Learn bsearch.")
}

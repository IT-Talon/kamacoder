package main

import (
	"fmt"
)

func main() {
	var a [4]int
	c := cap(a)
	fmt.Println(c)
	fmt.Println(7 >> 2)
	fmt.Println(quicksort3([]int{2, 56, 32, 6, 2, 1}))
	fmt.Println(quicksort2([]int{2, 56, 32, 6, 2, 1}))
	fmt.Println(quicksort([]int{2, 56, 32, 6, 2, 1}))
}

// 分治，大的往右堆，小的往左堆，然后大小堆继续递归排序，最后合并
func quicksort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	pivot := data[0]
	var small, large, equal []int
	for _, datum := range data {
		if datum < pivot {
			small = append(small, datum)
		} else if datum == pivot {
			equal = append(equal, datum)
		} else {
			large = append(large, datum)
		}
	}
	small = quicksort(small)
	large = quicksort(large)
	return append(append(small, equal...), large...)
}

// 2, 56, 32, 6, 2, 1
// 分治，遍历数组，大的往右堆，小的往左堆，将选择的轴点交换到大小堆的中心点来，然后递归排序中心点右边的大堆和中心点左边的小堆
func quicksort2(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	pivot := data[0]
	i := 1
	for j := 1; j < len(data); j++ {
		if data[j] < pivot {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[0], data[i-1] = data[i-1], data[0]
	quicksort2(data[:i])
	quicksort2(data[i:])
	return data
}

// 2, 56, 32, 6, 2, 1
// 分治，遍历数组，大的往右堆，小的往左堆，将选择的轴点交换到大小堆的中心点来，然后递归排序中心点右边的大堆和中心点左边的小堆
func quicksort3(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	pivot := data[0]
	i, j := 1, len(data)-1
	for i < j {
		if data[i] > pivot {
			data[i], data[j] = data[j], data[i]
			j--
		} else {
			i++
		}
	}
	if data[i] < pivot {
		data[i], data[0] = data[0], data[i]
	}
	quicksort3(data[:i])
	quicksort3(data[i:])
	return data
}

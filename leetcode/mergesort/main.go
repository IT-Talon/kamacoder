package main

import "fmt"

func main() {
	var data = []int{2, 56, 32, 6, 2, 1}
	fmt.Println(mergeSort(data))
}

/*
*
2 56 32 6 2 1

2 56 32   6 2 1
2    56 32     6   2 1
*/
func mergeSort(data []int) []int {
	if len(data) > 1 {
		data1 := mergeSort(data[:len(data)/2])
		data2 := mergeSort(data[len(data)/2:])
		return merge(data1, data2)
	}
	return data

}

func merge(data1, data2 []int) []int {

	var res []int
	var i, j int
	for i < len(data1) || j < len(data2) {
		if i < len(data1) && j < len(data2) {
			if data1[i] < data2[j] {
				res = append(res, data1[i])
				i++
			} else {
				res = append(res, data2[j])
				j++
			}
			continue
		}
		if i < len(data1) {
			res = append(res, data1[i])
			i++
		}
		if j < len(data2) {
			res = append(res, data2[j])
			j++
		}
	}
	return res
}

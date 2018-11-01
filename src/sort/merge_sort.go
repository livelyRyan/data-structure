package main

import "fmt"

// 归并排序
// mergeSort(p, q)  = merge ( mergeSort(p,r) , mergeSort(r+1, q) )
// 终止条件  p >= q
func main() {

	arr := []int{84, 3, 20, 34, 94, 19, 9, 438, 178}

	mergeSort(&arr, 0, len(arr)-1)

	PrintArr(arr)
}

func mergeSort(arr *[]int, start, end int) {

	if start >= end {
		return
	}

	mid := (start + end) / 2

	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)

	merge(arr, start, mid, end)
}

func merge(arr *[]int, start, mid, end int) {

	arr1 := make([]int, mid - start + 1)
	copy(arr1, (*arr)[ start : mid + 1])

	arr2 := make([]int, end - mid )
	copy(arr2, (*arr)[ start : mid + 1])

	//
	arr1 = append(arr1, (*arr)[mid] + (*arr)[end])
	arr2 = append(arr2, (*arr)[mid] + (*arr)[end])


	i := 0
	j := 0
	for ; start <= end ; start++ {
		if arr1[i] < arr2[j] {
			(*arr)[start] = arr1[i]
			i++
		} else {
			(*arr)[start] = arr2[j]
			j++
		}

	}
}


func PrintArr(arr []int) {
	fmt.Print("[")
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])

		if i != len(arr) - 1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")

}

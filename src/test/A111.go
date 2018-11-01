package main

import "fmt"

func main() {
	arr := []int{84, 3, 20}
	a1 := arr[0:3]
	a1[0] = 100
	fmt.Print(arr)
}

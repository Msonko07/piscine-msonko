package main

import "fmt"

func NumInList(list []int, num int) bool {
	for _, i := range list {
		if i == num {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(NumInList([]int{4, 4, 2, 5}, 5))
	fmt.Println(NumInList([]int{5, 5, 5, 5}, 5))
	fmt.Println(NumInList([]int{4, -2, 0, 3}, 5))
}

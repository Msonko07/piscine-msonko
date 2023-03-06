package main

import "fmt"

func SumList(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum = sum + i
	}
	return sum

}

func main() {
	fmt.Println(SumList([]int{4, 4, 2, 4}))  // 14
	fmt.Println(SumList([]int{5, 5, 5, 5}))  // 20
	fmt.Println(SumList([]int{4, -2, 0, 3})) // 5
	fmt.Println(SumList(nil))                // 0
	fmt.Println(SumList([]int{}))            // 0

}

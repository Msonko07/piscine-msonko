package main

import "fmt"

func tabMult(n int) {
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d x %d = %d\n ", i, n, i*n)
	}
}

func main() {
	tabMult(9)
	fmt.Println()
}

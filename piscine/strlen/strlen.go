package main

import "fmt"

func StrLen(s string) int {
	for i := 0; i <= len(s); i++ {
		fmt.Print()
	}
	return len(s)
}

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}

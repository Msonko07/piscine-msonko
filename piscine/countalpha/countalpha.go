package main

import "fmt"

func CountAlpha(s string) int {
	runes := []rune(s)
	count := 0
	for _, letter := range runes {
		if letter >= 'a' && letter <= 'z' || letter >= 'A' && letter <= 'Z' {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountAlpha("Hello World"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))

}

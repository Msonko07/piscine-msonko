package main

import "github.com/01-edu/z01"

func LastRune(s string) rune {
	sequence_rune := len([]rune(s)) - 1
	return []rune(s)[sequence_rune]

}

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}

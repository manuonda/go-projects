package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "abcd"

	fmt.Println("Len : ", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}

	fmt.Println()

	fmt.Println("Rune count : ", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf(" %s starts at %d", runeValue, idx)
	}

	fmt.Println("Using DecodeRuneInString")

}

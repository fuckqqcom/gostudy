package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes你好"
	fmt.Println(s)
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}

	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		fmt.Printf("chis %c ", ch)

		bytes = bytes[size:]
		fmt.Println(size, bytes)

	}

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}

}

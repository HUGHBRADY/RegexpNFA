package main

import (
	"fmt"
)

func intopost(infix string) string {
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

	postfix, s := []rune{}, []rune{}

	return string(postfix)
}

func main() {
	// Answer ab.c*
	fmt.Println("Infix:		", "a.b.c*")
	fmt.Println("Postfix: 	", intopost("a.b.c*"))

	// Answer abd|.*
	fmt.Println("Infix:		", "(a.(b|d).c*")
	fmt.Println("Postfix: 	", intopost("(a.(b|d).c*"))

	// Answer abd|.c*.
	fmt.Println("Infix:		", "a.(b|d).c*")
	fmt.Println("Postfix: 	", intopost("a.(b|d).c*"))

	// Answer abb.+.c
	fmt.Println("Infix:		", "a.(b.b)+.c")
	fmt.Println("Postfix: 	", intopost("a.(b.b)+.c"))
}
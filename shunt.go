package main

import (
	"fmt"
)

func intopost(infix string) string {
	// Maps runes to ints. These are ordered by precedence in regexps
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

	postfix, s := []rune{}, []rune{}

	// Loop over infix string a char at a time
	for _, r := range infix {
		switch {
		case r == '(':
			// Throws open bracket onto s stack
			s = append(s, r)
		case r == ')':
			// Loop through stack until you see open bracket
			for s[len(s)-1] != '(' {
				// Keep popping chars off stack and appending to postfix
				postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
			}
			// Finally, discard open bracket
			s = s[:len(s)-1]
		case specials[r] > 0:
			// while stack has elements, and precedence of the character <= the precedence of last element
			for len(s) > 0 && specials [r] <= specials[s[len(s)-1]] {
				// Pop the elements off top of stack and append to postfix
				postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)
		// Normal characters eg a, b, c
		default:
			// Appends r to the end of the postfix[]
			postfix = append(postfix, r)
		}
	}

	// If anything is left in the stack pop it into postfix 
	for len(s) > 0 {
		postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
	}

	return string(postfix)
}

func main() {
	// Answer ab.c*
	fmt.Println("Infix:		", "a.b|c*")
	fmt.Println("Postfix: 	", intopost("a.b|c*"))

	// Answer abd|.*
	fmt.Println("Infix:		", "(a.(b|d))")
	fmt.Println("Postfix: 	", intopost("(a.(b|d))"))

	// Answer abd|.c*.
	fmt.Println("Infix:		", "a.(b|d).c*")
	fmt.Println("Postfix: 	", intopost("a.(b|d).c*"))

	// Answer abb.+.c
	fmt.Println("Infix:		", "a.(b.b)+.c")
	fmt.Println("Postfix: 	", intopost("a.(b.b)+.c"))
}
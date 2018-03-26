package main

import (
	"fmt"
)

type state struct {
	symbol rune 
	edge1 *state
	edge2 *state
}

type nfafrag struct {
	initial *state
	accept  *state
}

func poregtonfa(postfix string) *nfa {
	// Provides an array of pointers to nfa (struct above) that is empty
	nfastack := []*nfa{}	
	
		for _, r := pofix {
			switch r {
			case '.':
				
			case '|':

			case '*':

			default:

			}
		}

	// DO ERROR CHECKING TO MAKE SURE THERE'S ONLY ONE ELEMENT ON NFA STACK
	return nfastack[0]
}

func main() {
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}
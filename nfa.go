package main

import (
	"fmt"
)

type state struct {
	symbol rune 
	edge1 *state
	edge2 *state
}

type nfa struct {
	initial *state
	accept  *state
}

func poregtonfa(postfix string) *nfa {
	// Provides an array of pointers to nfa (struct above) that is empty
	nfastack := []*nfa{}	
	
		for _, r := range postfix {
			switch r {
			case '.':
				// Take two elements off nfa stack
				frag2 := nfastack[len(nfastack)-1]
				// Get rid of last element on stack
				nfastack = nfastack[:len(nfastack)-1]

				frag1 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				// Joins the accept state of frag1 to initial state of frag2
				frag1.accept.edge1 = frag2.initial

				nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})

			case '|':
				// Take two elements off nfa stack
				frag2 := nfastack[len(nfastack)-1]
				// Get rid of last element on stack
				nfastack = nfastack[:len(nfastack)-1]

				frag1 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				accept := state{}
				initial := state{edge1: frag1.initial, edge2: frag2.initial}
				frag1.accept.edge1 = &accept
				frag2.accept.edge1 = &accept
				 
				nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			case '*':
				// Pop a fragment off the stack
				frag := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				accept := state{}
				initial := state{edge1: frag.initial, edge2: &accept}
				// New fragment is old frag with 2 extra states (new initial and accept)
				frag.accept.edge1 = frag.initial
				frag.accept.edge2 = &accept

				// Push new fragment to nfa stack
				nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			default:
				accept := state{}
				initial := state{symbol: r, edge1: &accept}

				nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			}
		}

	// DO ERROR CHECKING TO MAKE SURE THERE'S ONLY ONE ELEMENT ON NFA STACK
	return nfastack[0]
}

func main() {
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}
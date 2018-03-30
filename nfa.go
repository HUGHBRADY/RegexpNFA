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

// Will add a state as well as all states that can then be reached using empty strings 
func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)

	// If state has empty strings "paths"
	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}

	return l
}

// Function that checks if regexp matches string
func postmatch(po string, input string) bool {
	ismatch := false
	ponfa := poregtonfa(po)

	// List of states you're currently in in NFA
	current := []*state{}
	next := []*state{}

	// Initialise current array with initial states in NFA 
	current = addState(current[:], ponfa.initial, ponfa.accept)

	// Loop through input string
	for _, r := range input {
		// Loop through current state
		for _, s := range current {
			// If symbol in current state == the character you're reading from input
			if s.symbol == r {
				next = addState(next[:], s.edge1, ponfa.accept)
			}
		}
		// Move from current state to the next state
		current, next = next, []*state{}
	}

	for _, s := range current {
		// If current (final) state is accepted
		if s == ponfa.accept {
			ismatch = true
			break
		}
	}

	return ismatch
}

func main() {
	fmt.Println(postmatch("ab.c*|", "def"))
}
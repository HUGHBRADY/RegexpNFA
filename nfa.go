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

// This function transforms regular expressions from infix to postfix.
func topostfix(infix string) string {
	// Maps runes to ints. These are ordered by precedence.
	specials := map[rune]int{'*': 10, '.': 9, '|': 8, '+': 11, '?': 12}

	postfix, s := []rune{}, []rune{}

	// Loop over infix string a char at a time.
	for _, r := range infix {
		switch {
		case r == '(':
			// Throws open bracket onto stack.
			s = append(s, r)
		case r == ')':
			// Loop through stack until you see open bracket.
			for s[len(s)-1] != '(' {
				// Keep popping chars off stack and appending to postfix.
				postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
			}
			// Finally, discard open bracket.
			s = s[:len(s)-1]
		case specials[r] > 0:
			// while stack has elements, and precedence of the character <= the precedence of last element.
			for len(s) > 0 && specials [r] <= specials[s[len(s)-1]] {
				// Pop the elements off top of stack and append to postfix.
				postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)
		// Normal characters eg a, b, c
		default:
			// Appends r to the end of the postfix[]
			postfix = append(postfix, r)
		}
	}

	// If anything is left in the stack pop it into postfix.
	for len(s) > 0 {
		postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
	}

	return string(postfix)
}

func poregtonfa(postfix string) *nfa {
	// Provides an array of pointers to nfa (struct above) that is empty.
	nfastack := []*nfa{}	
	
		for _, r := range postfix {
			switch r {
			case '.':
				// Take two fragments off nfa stack and get rid of last element on stack.
				frag2 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				frag1 := nfastack[len(nfastack)-1]				
				nfastack = nfastack[:len(nfastack)-1]

				// Joins the accept state of frag1 to initial state of frag2.
				frag1.accept.edge1 = frag2.initial

				nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})

			case '|':
				// Take two fragments off nfa stack and get rid of last element on stack.
				frag2 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				frag1 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				accept := state{}
				initial := state{edge1: frag1.initial, edge2: frag2.initial}
				frag1.accept.edge1 = &accept
				frag2.accept.edge1 = &accept
				 
				nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			case '*':
				// Pop a fragment off the stack.
				frag := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				accept := state{}
				initial := state{edge1: frag.initial, edge2: &accept}
				// Define where accept state of fragment points.
				frag.accept.edge1 = frag.initial
				frag.accept.edge2 = &accept

				// Push new fragment to nfa stack.
				nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})	
			case '+':
				// Pop a fragment off the stack.
				frag := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				accept := state{}
				initial := state{edge1: frag.initial}
				middle := state{edge1: frag.initial, edge2: &accept}

				frag.accept.edge1 = &middle

				// Push new fragment to nfa stack.
				nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			case '?':
				// Pop a fragment off the stack.
				frag := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				accept := state{}
				initial := state{edge1: frag.initial, edge2: &accept}

				frag.accept.edge1 = &accept

				// Push new fragment to NFA stack.
				nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			default:
				accept := state{}
				initial := state{symbol: r, edge1: &accept}

				nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			}
		}

	// Make sure only one element in nfaStack.
	if len(nfastack) != 1 {
		fmt.Println("Error! NFA invalid ", len(nfastack), nfastack)
	}

	return nfastack[0]
}

// Will add a state as well as all states that can then be reached using empty strings.
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

// Function that checks if regexp matches string.
func regexpmatch(userexp string, input string) bool {
	
	ismatch := false
	ponfa := poregtonfa(userexp)

	// List of states you're currently in in NFA.
	current := []*state{}
	next := []*state{}

	// Initialise current array with initial states in NFA.
	current = addState(current[:], ponfa.initial, ponfa.accept)

	// Loop through input string.
	for _, r := range input {
		// Loop through current state.
		for _, s := range current {
			// If symbol in current state == the character you're reading from input.
			if s.symbol == r {
				next = addState(next[:], s.edge1, ponfa.accept)
			}
		}
		// Move from current state to the next state.
		current, next = next, []*state{}
	}

	for _, s := range current {
		// If current (final) state is accepted.
		if s == ponfa.accept {
			ismatch = true
			break
		}
	}

	return ismatch
}

func main() {
	var userexp, poreg, userstring string

	fmt.Print("Enter a regexp: ")
	fmt.Scanln(&userexp)

	// Convert expression into postfix.
	poreg = topostfix(userexp)
	fmt.Println("\nPostfix: ", poreg)

	fmt.Print("Enter a string to test: ")
	fmt.Scan(&userstring)

	// Check if string matches regexp.
	fmt.Println(regexpmatch(poreg, userstring))
}
# Graph Theory Project 2018
# RegexpNFA
A program written in Go that can build a non-deterministic finite automata (NFA) from a regular expression, that can use the NFA to check if the regular expression matches any given string of text.

# How to run
Before you can run this program you must have Go installed on your machine. [Link to download.](https://golang.org/dl/) 

Once Go is installed and you have downloaded this repository onto your machine, you can run the program. Open the command prompt and navigate to the program folder. Type "go run nfa.go" and hit enter. You can also run the program through Visual Studio Code's terminal.

On startup, you will be prompted to enter a regular expression (which must be in infix notation). Next you will be prompted to enter a string to test against it. The program will then tell you if the string matched the regular expression by way of true or false statements.

# How it works
This project was designed in order to test our ability to put theory into practice. This project employs two algorithms; Thompson's construction and the shunting-yard algorithm. Thompson's construction is an algorithm for creating NFAs from regular expressions. It employs two structs; state and nfa, that are used to represent the NFA. As the program reads in the regexp, it breaks up the NFA and reassembles it with the added parameters. 

Once the NFA is finished, the program tests each character from the input string. This is done by looping through both the input string and the states in the NFA and checking if the characters match. If the program ends up in the accept state, the string is accepted.

The shunting-yard algorithm is used to transform regular expressions from infix to postfix notation. This is necessary for the program to understand the regular expressions thanks to the lack of parentheses.

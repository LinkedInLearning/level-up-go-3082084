package main

import (
	"flag"
	"log"
)

// operatorType represents the type of operator in an expression
type operatorType int

const (
	openBracket operatorType = iota
	closedBracket
	otherOperator
)

// bracketPairs is the map legal bracket pairs
var bracketPairs = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

// getOperatorType return the operator type of the give operator.
func getOperatorType(op rune) operatorType {
	for ob, cb := range bracketPairs {
		switch op {
		case ob:
			return openBracket
		case cb:
			return closedBracket
		}
	}

	return otherOperator
}

// stack is a simple LIFO stack implementation using a slice.
type stack struct {
	elems []rune
}

// push adds a new element to the stack.
func (s *stack) push(e rune) {
	s.elems = append(s.elems, e)
}

// pop removes the last element from the stack.
func (s *stack) pop() *rune {
	if len(s.elems) == 0 {
		return nil
	}
	n := len(s.elems) - 1
	last := s.elems[n]
	s.elems = s.elems[:n]
	return &last
}

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	s := stack{}
	for _, e := range expr {
		switch getOperatorType(e) {
		case openBracket:
			s.push(e)
		case closedBracket:
			last := s.pop()
			if last == nil || bracketPairs[*last] != e {
				return false
			}
		}
	}

	return len(s.elems) == 0
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

type Stack interface {
	push(r rune)
	pop() rune
}

type RuneStack struct {
	Collection []rune
	length     int
}

func (rs *RuneStack) push(r rune) {
	rs.Collection[rs.length] = r
	rs.length++
}

func (rs *RuneStack) pop() rune {
	last := rs.length - 1
	lastItem := rs.Collection[last]
	rs.Collection = append(rs.Collection[:last], rs.Collection[last+1:]...)
	rs.length--
	return lastItem
}

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	brackets := 
	rs := RuneStack{Collection: make([]rune, 0), length: 0} 

	runes := strings.Map(func(r rune) rune {
		return r
	}, expr)

	for _, r := range runes {}
	fmt.Println(runes)

	return true
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

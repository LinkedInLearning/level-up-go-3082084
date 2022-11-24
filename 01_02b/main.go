package main

import (
	"bytes"
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	wordSlice := strings.Split(msg, " ")
	// fmt.Println(singleLetter)


	for _, word := range wordSlice {
		var buffer bytes.Buffer
		letterSlice := strings.Split(word, "")
		for i, letter := range letterSlice{
			for j:=0; j < i+1; j++{
				buffer.WriteString(letter)
			}
		}
		print(buffer.String())
	}

}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}

package main

import (
	"flag"
	"log"
)

var messages = []string{
	"Hello!",
	"How are you?",
	"Are you just going to repeat what I say?",
	"So immature",
	"Stop copying me!",
}

// repeat concurrently prints out the given message n times
func repeat(n int, message string) {
	ch := make(chan struct{})
	for i := 0; i < n; i++ {
		go func(i int) {
			log.Printf("[G%d]:%s\n", i, message)
			ch <- struct{}{}
		}(i)
	}

	for i := 0; i < n; i++ {
		<-ch
	}
}

func main() {
	factor := flag.Int64("factor", 0, "The fan-out factor to repeat by")
	flag.Parse()

	for _, m := range messages {
		log.Printf("[Main]:%s\n", m)
		repeat(int(*factor), m)
	}
}

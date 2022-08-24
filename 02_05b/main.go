package main

import (
	"fmt"
	"log"
)

// setup constants
const bartenderCount = 3
const partyGoerCount = 20
const maxDrinkCount = 40

// the total amount of drinks that the bartenders have made
type party struct {
	drinkCount int

	orderDrink  chan struct{}
	finishDrink chan struct{}
}

// registerDrink ensures that the drink made by the bartenders is counted.
func (p *party) registerDrink() {
	p.drinkCount++
}

// bartender is the bartender functionality of the party
func (p *party) bartender(name string) {
	for {
		select {
		case <-p.orderDrink:
			p.registerDrink()
			log.Printf("%s makes a drink.\n", name)
			p.finishDrink <- struct{}{}
		}
	}
}

// partyGoer is the partygoer functionality of the party
func (p *party) partyGoer(name string) {
	for {
		select {
		case p.orderDrink <- struct{}{}:
			log.Printf("%s orders a drink!", name)
			<-p.finishDrink
			log.Printf("%s enjoys a drink!\n", name)
		}
	}
}

func main() {
	log.Println("Welcome to the Level Up Go party!")
	orderDrink := make(chan struct{}, bartenderCount)
	finishDrink := make(chan struct{}, bartenderCount)
	p := party{
		orderDrink:  orderDrink,
		finishDrink: finishDrink,
	}
	for i := 0; i < bartenderCount; i++ {
		go p.bartender(fmt.Sprint("Bartender-", i))
	}
	for i := 0; i < partyGoerCount; i++ {
		go p.partyGoer(fmt.Sprint("Partygoer-", i))
	}
	log.Println("The Level Up Go party has ended! Good night!")
}

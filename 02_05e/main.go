package main

import (
	"fmt"
	"log"
	"sync"
)

// setup constants
const bartenderCount = 3
const partyGoerCount = 20
const maxDrinkCount = 40

// the total amount of drinks that the bartenders have made
type party struct {
	drinkCount int
	drinksLock sync.Mutex

	orderDrink  chan struct{}
	finishDrink chan struct{}
	endParty    chan struct{}
}

// registerDrink ensures that the drink made by the bartenders is incremented.
func (p *party) registerDrink() {
	p.drinksLock.Lock()
	defer p.drinksLock.Unlock()
	p.drinkCount++
	if p.drinkCount == maxDrinkCount {
		close(p.endParty)
		return
	}
}

// bartender is the bartender functionality of the party
func (p *party) bartender(name string) {
	for {
		select {
		case <-p.orderDrink:
			p.registerDrink()
			log.Printf("%s makes a drink.\n", name)
			p.finishDrink <- struct{}{}
		case <-p.endParty:
			log.Printf("%s goes home. Bye!\n", name)
			return
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
		case <-p.endParty:
			log.Printf("%s goes home! Bye!\n", name)
			return
		}
	}
}

func main() {
	log.Println("Welcome to the Level Up Go party!")
	orderDrink := make(chan struct{}, bartenderCount)
	finishDrink := make(chan struct{}, bartenderCount)
	endParty := make(chan struct{})
	p := party{
		orderDrink:  orderDrink,
		finishDrink: finishDrink,
		endParty:    endParty,
	}
	for i := 0; i < bartenderCount; i++ {
		go p.bartender(fmt.Sprint("Bartender-", i))
	}
	for i := 0; i < partyGoerCount; i++ {
		go p.partyGoer(fmt.Sprint("Partygoer-", i))
	}
	<-endParty
	log.Println("The Level Up Go party has ended! Good night!")
}

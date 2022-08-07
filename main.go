package main

import (
	"log"
	"math/rand"
	"time"
)

const path = "entries.json"

// raffleEntries is the struct we unmarshal raffle entries into
type raffleEntries struct {
	// TODO: Fill in definition
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntries {
	panic("NOT IMPLEMENTED")
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntries) raffleEntries {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}
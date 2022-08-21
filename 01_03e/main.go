package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "entries.json"

// raffleEntries is the struct we unmarshal raffle entries into
type raffleEntries struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntries {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []raffleEntries
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
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

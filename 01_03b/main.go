package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

const path = "entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	id   int
	name string
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error:",err)
	}
	// fmt.Println(file)
	strFile := string(file)
	// println(strFile)

	var entries []raffleEntry

	error := json.Unmarshal([]byte(strFile), &entries)
	if error != nil {
		log.Fatal("Error occured:", error)
	}
	fmt.Println(entries)
	return entries

}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
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

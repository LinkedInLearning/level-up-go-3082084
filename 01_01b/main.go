package main

import (
	"fmt"
	"log"
	"time"
)
var expectedFormat = "2006-01-02"
// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	inputTime, err := time.Parse(expectedFormat, target)
	if err != nil || time.Now().After(inputTime) {
		log.Fatal("invalid target date", target)
	}
	return inputTime
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	sleepTime := time.Until(target)
	return sleepTime.Hours() / 24
}

func main() {
	var bday string
	fmt.Println("Your next bday in YYYY-MM-DD format")
	fmt.Scan(&bday)

	target := parseTime(bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!",
	int(calcSleeps(target)))
}

package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

var (
	showYear bool
)

func main() {
	flag.BoolVar(&showYear, "y", false, "day-progress -y")

	flag.Parse()

	printDayProgress()

	if showYear {
		printYearProgress()
	}
}

func printYearProgress() {
	dayOfYear := time.Now().YearDay()
	daysInYear := 365

	percentage := int(math.Floor(float64(dayOfYear) / float64(daysInYear) * 100))

	fmt.Print("Year: [")

	for i := 0; i < 100; i += 10 {
		fmt.Print("#")
	}

	fmt.Printf("] %d%%\n", percentage)
}

func printDayProgress() {
	timeWhenDayStarted := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Local().Location()).Unix()
	now := time.Now().Unix()
	offset := float64(now - timeWhenDayStarted)
	secondsInDay := 86400
	percentage := int(math.Round(offset / float64(secondsInDay) * 100))

	if showYear {
		fmt.Print("Day:  [")
	} else {
		fmt.Print("Day: [")
	}

	for i := 0; i < 100; i += 10 {
		if i < percentage {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
	}

	fmt.Printf("] %d%%\n", percentage)
}

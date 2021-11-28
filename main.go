package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	printDayProgress()
}

func printDayProgress() {
	timeWhenDayStarted := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Local().Location()).Unix()
	now := time.Now().Unix()
	offset := float64(now - timeWhenDayStarted)
	secondsInDay := 86400
	percentage := int(math.Round(offset / float64(secondsInDay) * 100))

	fmt.Print("Day: [")

	for i := 0; i <= 100; i += 10 {
		if i < percentage {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
	}

	fmt.Printf("] %d%%\n", percentage)
}

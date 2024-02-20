package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	hour := currentTime.Hour()

	var welcomeMessage string

	var timeSelect int

	if hour < 7 || hour > 23 {
		timeSelect = 0
	}

	if hour >= 7 && hour < 12 {
		timeSelect = 1
	}

	if hour >= 12 && hour < 18 {
		timeSelect = 2
	}

	if hour >= 18 && hour < 23 {
		timeSelect = 3
	}

	switch timeSelect {
	case 0:
		welcomeMessage = "Sorry, de parkeerplaats is ’s nachts gesloten"
	case 1:
		welcomeMessage = "Goedemorgen "
	case 2:
		welcomeMessage = "Goedemiddag "
	case 3:
		welcomeMessage = "Goedenavond "
	default:
		fmt.Println("An error has occurred please contact customer support at 1234 not happening")
	}

	fmt.Print(welcomeMessage)
	if timeSelect != (0) {
		fmt.Println("Welkom bij Fonteyn Vakantieparken")
	}
}

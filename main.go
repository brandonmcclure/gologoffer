package main

import (
	"log"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	var startHour, endHour int
	var isDryRun bool
	
	switch len(os.Args) {
	case 5:
		isDryRun = os.Args[4] == "--dry-run"
		fallthrough
	case 4:
		startHour = parseHour(os.Args[2])
		endHour = parseHour(os.Args[3])
	case 3:
		startHour = parseHour(os.Args[2])
		endHour = 6 // Default end hour if only start hour is provided
	default:
		startHour = 22 // Default start hour
		endHour = 6   // Default end hour
	}

	currentTime := time.Now()

	if currentTime.Hour() >= startHour || currentTime.Hour() < endHour {
		if isUserLoggedIn() {
			displayNotification("Computer Shutdown", "The computer will be shut down soon.")
			if !isDryRun {
				initiateShutdown()
			}
		}
	} else {
		fmt.Print("Outside of time range, no action taken")
	}
}
func parseHour(hourArg string) int {
	hour, err := strconv.Atoi(hourArg)
	if err != nil {
		log.Fatalf("Invalid hour value: %s", err)
	}
	if hour < 0 || hour > 23 {
		log.Fatal("Hour value must be between 0 and 23")
	}
	return hour
}

func initiateShutdown() {
	cmd := exec.Command("shutdown", "-h", "now")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func isUserLoggedIn() bool {
	cmd := exec.Command("who")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	return len(strings.TrimSpace(string(output))) > 0
}

func displayNotification(title, message string) {
	cmd := exec.Command("notify-send", title, message)
	cmd.Env = append(os.Environ(), "DISPLAY=:0")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

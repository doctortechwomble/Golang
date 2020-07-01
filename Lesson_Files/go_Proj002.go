package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Setting up variable
	rand.Seed(time.Now().UnixNano())

	isHeistOn := true

	//First Condition (Sneak past guards)
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good  job, but remember, this is the first step.\n")
	} else {
		fmt.Println("Plan a better disguise next time!\n")
	}

	//Second Condition (Open the vault)
	openedVault := rand.Intn(100)

	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!\n")
	} else {
		fmt.Println("What's the combo to this lock again??\n")
	}

	//Third Condition (Leaving)
	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Looks like you tripped an alarm\n")
		case 1:
			isHeistOn = false
			fmt.Println("Vault doors jammed\n")
		case 2:
			isHeistOn = false
			fmt.Println("You were beaten by an angry old dear\n")
		case 3:
			isHeistOn = false
			fmt.Println("Superman came to save the day\n")
		case 4:
			isHeistOn = false
			fmt.Println("Bank doors locked")
		default:
			fmt.Println("Start the getaway car\n")
		}
	}
	//Fourth Condition (Amount Stolen)
	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("We got away with: £", amtStolen, "\n")
	}

	fmt.Println("Is the Heist currently proceeding?", isHeistOn)
}

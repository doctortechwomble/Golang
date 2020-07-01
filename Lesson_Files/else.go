package main

import "fmt"

func main() {
	lockCombo := "2-35-19"
	robAttempt := "2-35-19"

  // Add your code below:
  if (lockCombo == robAttempt) {
    fmt.Println("The vault is now opened.")
  } else {
    fmt.Println("Piss Off!")
  }
}

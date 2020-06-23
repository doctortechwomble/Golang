package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("Hello what's you name?\n")
	fmt.Scan(&name)
	fmt.Printf("\nHi %v, great to meet you, I'm Go.\n \nHow old are you?\n", name)
	fmt.Scan(&age)
	fmt.Printf("\nWow as old as %d, I'm only 9!\n", age)
}

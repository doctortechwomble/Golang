package main

import "fmt"

func main() {
	var publisher, writer, artist, title string
	var year, pagenumber uint16
	var grade float32
	var pricenew, salesprice float32

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pagenumber = 14
	grade = 6.5
	pricenew = 20.99
	salesprice = grade / 100

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "In year of our Lord", year, "The dirty pics are found on", pagenumber, "The condition of item is", grade, "\n")

	fmt.Println("Our bargain price £", grade/100*pricenew, "\n")

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pagenumber = 160
	grade = 9.0
	pricenew = 12.32
	salesprice = grade / 100

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "In year of our Lord", year, "The dirty pics are found on", pagenumber, "The condition of item is", grade, "Our sales price", salesprice, "\n")

	fmt.Println("Our bargain price £", grade/100*pricenew, "\n")

}

package main

import (
	"fmt"
)

func main() {
	var publisher, writer, artist, title string
	var year, pageNumber uint
	var grade float32
	title = "Mr.GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5
	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "published in", year, "pages", pageNumber, "condition grade", grade)
}
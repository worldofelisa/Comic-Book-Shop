package main

import (
	"fmt"
)

type Book struct {
	Title     string
	Author    string
	Artist    string
	Publisher string
	Year      int
	Pages     int
	Condition float64
}

var bookList = []Book{
	{
		Title:     "Mr.GoToSleep",
		Author:    "Tracey Hatchet",
		Artist:    "Jewel Tampson",
		Publisher: "DizzyBooks Publishing Inc.",
		Year:      1997,
		Pages:     14,
		Condition: 6.5,
	},
	{
		Title:     "Epic Vol. 1",
		Author:    "Ryan N. Shawn",
		Artist:    "Pheobe Paperclips",
		Publisher: "DizzyBooks Publishing Inc.",
		Year:      2013,
		Pages:     160,
		Condition: 9.0,
	},
	{
		Title:     "The Adventures of Mr. Man",
		Author:    "Lord Shellington",
		Artist:    "Marco Doli",
		Publisher: "Printing Book Press",
		Year:      2016,
		Pages:     55,
		Condition: 8.7,
	},
}

func main() {
	for _, book := range bookList {
		fmt.Printf("Title %s\n Written by %s\n Drawn by %s\n Published by %s\n Published in %d\n With %d Pages \n In %.2f Condition \n \n", book.Title, book.Author, book.Artist, book.Publisher, book.Year, book.Pages, book.Condition)
	}
}

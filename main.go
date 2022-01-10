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

//TODO If you were to add data from other points or allow others to add data,
//add a validation here to stop people from breaking stuff

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
	{
		Title:     "Mia's \"Weird\" Day Out",
		Author:    "Mia McKelly",
		Artist:    "Kelly Miana",
		Publisher: "Day Adventures LLC.",
		Year:      1995,
		Pages:     25,
		Condition: 8.25,
	},
}

func main() {
	for _, book := range bookList {
		fmt.Printf("Title %s\n Written by %s\n Drawn by %s\n Published by %s\n Published in %d\n With %d Pages \n In %.2f Condition \n \n", book.Title, book.Author, book.Artist, book.Publisher, book.Year, book.Pages, book.Condition)
	}
}

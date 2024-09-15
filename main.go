package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
}
type Book struct {
	Title         string
	Author        string
	Pages         int
	PublishedYear int
}

func main() {
	var car Car
	car.Brand = "Tesla"
	car.Model = "x"
	car.Year = 2021
	fmt.Println(car)

	book := Book{Title: "taras", Author: "taras", Pages: 23, PublishedYear: 77}

	fmt.Printf("Title %v Author %v Pages %v PublisherYear %v\n", book.Title, book.Author, book.Pages, book.PublishedYear)
}

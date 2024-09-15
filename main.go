package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
}

func (c Car) Greet() {
	fmt.Printf("Car brand is %v model %v year %v\n", c.Brand, c.Model, c.Year)
}

func (c *Car) ChangeYear() {
	c.Year++
}

type Book struct {
	Title         string
	Author        string
	Pages         int
	PublishedYear int
}

func (b Book) BookInfo() {
	fmt.Printf("Title %v Author %v Pages %v PublisherYear %v\n", b.Title, b.Author, b.Pages, b.PublishedYear)
}

func main() {
	car := Car{Brand: "Tesla", Model: "X", Year: 2024}
	car.Greet()
	car.ChangeYear()
	car.Greet()

	book := Book{Title: "taras", Author: "taras", Pages: 23, PublishedYear: 77}
	book.BookInfo()

}

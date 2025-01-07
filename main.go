package main

type status int

const (
	yetToRead status = iota
	currentlyReading
	completedReading
)

/* Custom Books*/

type Book struct {
	status      status
	title       string
	description string
}

//implement the book.item interface

func (t Book) FilterValue() string {
	return t.title
}

func (t Book) Title() string {
	return t.title
}

func (t Book) Description() string {
	return t.description
}

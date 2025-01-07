package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

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

/* MAIN MODEL*/
type Model struct {
	list list.Model
	err  error
}

func New() *Model{
	return &Model{}

}

func (m *Model) initList(width int, height int) {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.list.Title = "Book Tracker"
	m.list.SetItems([]list.Item{
		Book{
			status:      currentlyReading,
			title:       "abc",
			description: "xyz",
		},
		Book{
			status:      currentlyReading,
			title:       "pqr",
			description: "lmao",
		},
		Book{
			status:      currentlyReading,
			title:       "skibidi",
			description: "ohio",
		},
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.initList(msg.Width,msg.Height)
	} 


	var cmd tea.Cmd;
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.list.View();
}

func main (){
	m:= New();
	p:= tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

}
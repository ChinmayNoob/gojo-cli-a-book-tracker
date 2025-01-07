package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

const divisor = 4

const (
	yetToRead status = iota
	currentlyReading
	completedReading
)

/*Styling*/
var (
	columnStyle = lipgloss.
			NewStyle().Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#3c3c3c"))
	focusedStyle = lipgloss.
			NewStyle().Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
	helpStyle = lipgloss.
			NewStyle().Foreground(lipgloss.Color("241"))
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
	focused  status
	lists    []list.Model
	err      error
	loaded   bool
	quitting bool
}

func New() *Model {
	return &Model{}
}

// go to next list
func (m *Model) Next() {
	if m.focused == completedReading {
		m.focused = yetToRead
	} else {
		m.focused++
	}
}

// go to prev list
func (m *Model) Prev() {
	if m.focused == yetToRead {
		m.focused = completedReading
	} else {
		m.focused--
	}
}

func (m *Model) initLists(width int, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divisor, height/2)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}
	// init books accoring to their status

	//yet to read
	m.lists[yetToRead].Title = "Yet To Read"
	m.lists[yetToRead].SetItems([]list.Item{
		Book{
			status:      yetToRead,
			title:       "abc",
			description: "xyz",
		},
		Book{
			status:      yetToRead,
			title:       "pqr",
			description: "lmao",
		},
		Book{
			status:      yetToRead,
			title:       "skibidi",
			description: "ohio",
		},
	})
	//currently reading
	m.lists[currentlyReading].Title = "Currently Reading"
	m.lists[currentlyReading].SetItems([]list.Item{
		Book{
			status:      currentlyReading,
			title:       "abc",
			description: "xyz",
		},
	})
	//completed Reading
	m.lists[completedReading].Title = "Done Reading"
	m.lists[completedReading].SetItems([]list.Item{
		Book{
			status:      completedReading,
			title:       "abc",
			description: "xyz",
		},
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:

		if !m.loaded {
			columnStyle.Width(msg.Width/divisor)
			focusedStyle.Width(msg.Width/divisor)
			columnStyle.Height(msg.Height-divisor)
			focusedStyle.Height(msg.Height-divisor)
			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "left","h":
			m.Prev()
		case "right","l":
			m.Next()
		}
	}

	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	// If quitting, return an empty string.
	if m.quitting {
		return ""
	}

	// If not loaded yet, display a loading message.
	if !m.loaded {
		return "loading..."
	}

	// Extract the views for each column.
	yetToReadView := m.lists[yetToRead].View()
	currentlyReadingView := m.lists[currentlyReading].View()
	completedReadingView := m.lists[completedReading].View()

	// Switch based on the focused column to apply the `focusedStyle`.
	switch m.focused {
	case yetToRead:
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			focusedStyle.Render(yetToReadView),
			columnStyle.Render(currentlyReadingView),
			columnStyle.Render(completedReadingView),
		)
	case currentlyReading:
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			columnStyle.Render(yetToReadView),
			focusedStyle.Render(currentlyReadingView),
			columnStyle.Render(completedReadingView),
		)
	case completedReading:
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			columnStyle.Render(yetToReadView),
			columnStyle.Render(currentlyReadingView),
			focusedStyle.Render(completedReadingView),
		)
	default:
		// Fallback if `m.focused` is invalid.
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			columnStyle.Render(yetToReadView),
			columnStyle.Render(currentlyReadingView),
			columnStyle.Render(completedReadingView),
		)
	}
}


func main() {
	m := New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

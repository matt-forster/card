package main

import (
	"fmt"
	"os"
	"strings"

	cli "github.com/charmbracelet/bubbletea"
	tui "github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// Globals
type (
	model struct {
		name  string
		title string

		width int
	}
	theme struct {
		main      string
		secondary string
		highlight string
		subtle    string
		text      string
		error     string
	}

	adaptiveTheme struct {
		main      tui.AdaptiveColor
		secondary tui.AdaptiveColor
		highlight tui.AdaptiveColor
		subtle    tui.AdaptiveColor
		text      tui.AdaptiveColor
		error     tui.AdaptiveColor
	}
)

var (
	darkTheme = theme{
		main:      "#88c0d0",
		secondary: "#8fbcbb",
		highlight: "#a3be8c",
		subtle:    "#4c566a",
		text:      "#eceff4",
		error:     "#bf616a",
	}

	lightTheme = theme{
		main:      "#5e81ac",
		secondary: "#81a1c1",
		highlight: "#b48ead",
		text:      "#2e3440",
		subtle:    "#C0C0C0",
		error:     "#bf616a",
	}

	activeTheme = adaptiveTheme{
		main:      tui.AdaptiveColor{Light: lightTheme.main, Dark: darkTheme.main},
		secondary: tui.AdaptiveColor{Light: lightTheme.secondary, Dark: darkTheme.secondary},
		highlight: tui.AdaptiveColor{Light: lightTheme.highlight, Dark: darkTheme.highlight},
		subtle:    tui.AdaptiveColor{Light: lightTheme.subtle, Dark: darkTheme.subtle},
		text:      tui.AdaptiveColor{Light: lightTheme.text, Dark: darkTheme.text},
		error:     tui.AdaptiveColor{Light: lightTheme.error, Dark: darkTheme.error},
	}
)

var (
	docStyle = tui.NewStyle().
			Padding(1, 2, 1, 2).
			Margin(1, 1, 1, 1)

	contentBoxStyle = tui.NewStyle().
			Foreground(activeTheme.text).
			BorderForeground(activeTheme.main).
			Align(tui.Center).
			BorderStyle(tui.RoundedBorder()).
			BorderBottom(true).
			BorderLeft(true).
			BorderRight(true).
			PaddingTop(2)

	nameHeader = tui.NewStyle().
			Foreground(activeTheme.text).
			Align(tui.Center).
			Bold(true).
			BorderStyle(tui.RoundedBorder()).
			BorderForeground(activeTheme.highlight).
			PaddingLeft(2).
			PaddingRight(2)

	topBorder = tui.NewStyle().
			BorderForeground(activeTheme.main).
			BorderStyle(tui.RoundedBorder()).
			BorderTop(true)

	footerStyle = tui.NewStyle().
      Align(tui.Right).
			Padding(1).
			Foreground(activeTheme.subtle)
)

func initial() model {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	return model{
		name:  "Matt Forster",
		title: "Senior Software Engineer",

		width: width,
	}
}

func (m model) Init() cli.Cmd {
	return nil
}

func (m model) Update(msg cli.Msg) (cli.Model, cli.Cmd) {

	switch msg := msg.(type) {

	case cli.KeyMsg:
		switch msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, cli.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	doc := strings.Builder{}

	name := nameHeader.Render(m.name)
	leftTopBorder := topBorder.Copy().Width(m.width/2 - tui.Width(name)).BorderLeft(true)
	rightTopBorder := topBorder.Copy().Width(m.width/2 - tui.Width(name)).BorderRight(true)

	topContent := tui.JoinHorizontal(1, leftTopBorder.Render(""), name, rightTopBorder.Render(""))

	content := "hi"

	contentBox := contentBoxStyle.Width(tui.Width(topContent) - 2).Render(content)
	footerContent := footerStyle.Width(tui.Width(topContent) - 2).Render("(q): quit")

	doc.WriteString(tui.JoinVertical(0, topContent, contentBox, footerContent))
	return docStyle.MaxWidth(m.width).Render(doc.String())
}

func main() {

	program := cli.NewProgram(initial())

	if _, err := program.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

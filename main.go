package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	app := tea.NewProgram(
		NewApp(
			NewRoute("hello", NewText("Hello, world!")),
			NewRoute("goodbye", NewText("Goodbye, world!")),
		),
	)
	if _, err := app.Run(); err != nil {
		panic(err)
	}
}

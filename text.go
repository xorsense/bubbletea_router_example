package main

import tea "github.com/charmbracelet/bubbletea"

type Text struct {
	value string
}

func NewText(value string) Text {
	return Text{value: value}
}

func (t Text) Init() tea.Cmd {
	return nil
}

func (t Text) View() string {
	return t.value
}

func (t Text) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return t, nil
}

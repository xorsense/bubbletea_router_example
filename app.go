package main

import (
	"bytes"
	tea "github.com/charmbracelet/bubbletea"
	"text/template"
)

type Route struct {
	Key   string
	Value tea.Model
}

func NewRoute(key string, value tea.Model) Route {
	return Route{key, value}
}

type AppParams struct {
	IsNavigating bool
	CurrentRoute int
	Routes       []Route
	Child        string
}

type App struct {
	Child        tea.Model
	Routes       []Route
	currentRoute int
	isNavigating bool
}

func NewApp(routes ...Route) App {
	app := App{Routes: routes}
	if len(routes) > 0 {
		app.Child = routes[0].Value
	}
	return app
}

func (a App) Init() tea.Cmd {
	return tea.ClearScreen
}

func (a App) View() string {
	tmpl := `{{.Child}}

{{ if .IsNavigating -}}
	{{ $currentRoute := .CurrentRoute }}
	{{ range $i, $v := .Routes -}}
		{{ if eq $i $currentRoute }}→ {{ else }}  {{ end }}{{ $v.Key }}
	{{ end -}}
{{ end }}
      Exit: ctrl+c
  Navigate: ctrl+n
`
	params := AppParams{IsNavigating: a.isNavigating, CurrentRoute: a.currentRoute, Routes: a.Routes}

	if a.Child != nil {
		params.Child = a.Child.View()
	}

	t, _ := template.New("main").Parse(tmpl)
	var buf bytes.Buffer
	_ = t.Execute(&buf, params)

	return buf.String()
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return a, tea.Quit
		case tea.KeyCtrlN:
			a.isNavigating = !a.isNavigating
			return a, nil
		case tea.KeyDown, tea.KeyUp, tea.KeyEnter:
			if !a.isNavigating {
				break
			}
			switch msg.Type {
			case tea.KeyDown:
				if a.currentRoute < len(a.Routes)-1 {
					a.currentRoute += 1
				}
			case tea.KeyUp:
				if a.currentRoute > 0 {
					a.currentRoute -= 1
				}
			case tea.KeyEnter:
				a.Child = a.Routes[a.currentRoute].Value
			}
		}
	}
	if a.Child != nil {
		var cmd tea.Cmd
		a.Child, cmd = a.Child.Update(msg)
		return a, cmd
	}
	return a, nil
}

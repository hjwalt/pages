package component_sidebar_button_list

import (
	"embed"
	"html/template"

	"github.com/hjwalt/routes/example"
	"github.com/hjwalt/routes/page"
)

//go:embed *
var files embed.FS

var Html = template.Must(template.ParseFS(files, "component.html"))

type Model struct {
}

func New(m Model, items ...page.Component[example.Context]) page.Component[example.Context] {
	return page.NewComponentSlice[example.Context, Model](Html, m, items)
}

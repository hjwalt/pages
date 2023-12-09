package component_sidebar_item_header

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
	Label string
}

func New(m Model) page.Component[example.Context] {
	return page.NewComponentBasic[example.Context, Model](Html, m)
}

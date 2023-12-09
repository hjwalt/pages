package component_sidebar

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
	Icon   string
	Label  string
	Active bool
}

func New(m Model, top page.Component[example.Context], button page.Component[example.Context]) page.Component[example.Context] {
	items := map[string]page.Component[example.Context]{}

	items["top"] = top
	items["button"] = button

	return page.NewComponentMap[example.Context, Model](Html, m, items)
}

package component_sidebar_item

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/hjwalt/routes/example"
	"github.com/hjwalt/routes/page"
)

//go:embed *
var files embed.FS

var Html = template.Must(template.ParseFS(files, "component.html"))

type Model struct {
	Icon   string
	Label  string
	Link   string
	Active bool
}

func (c Model) Render(ctx example.Context, w http.ResponseWriter, r *http.Request) (template.HTML, error) {
	return page.ComponentRender[example.Context, Model](ctx, w, r, Html, c)
}

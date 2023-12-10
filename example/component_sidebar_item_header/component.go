package component_sidebar_item_header

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/hjwalt/routes/example"
	"github.com/hjwalt/routes/mvc"
)

//go:embed *
var files embed.FS

var Html = template.Must(template.ParseFS(files, "component.html"))

type Model struct {
	Label string
}

func (c Model) Render(ctx example.Context, w http.ResponseWriter, r *http.Request) (template.HTML, error) {
	return mvc.ComponentRender[example.Context, Model](ctx, w, r, Html, c)
}

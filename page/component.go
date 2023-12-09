package page

import (
	"context"
	"html/template"
	"net/http"
)

type Component[C context.Context] interface {
	Render(C, http.ResponseWriter, *http.Request) (template.HTML, error)
}

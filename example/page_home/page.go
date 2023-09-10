package page_home

import (
	"html/template"
	"net/http"

	"github.com/hjwalt/routes/example"
	"github.com/hjwalt/routes/example/page_error_500"
	"github.com/hjwalt/routes/runtime_chi"
	"github.com/hjwalt/runway/runtime"
)

const (
	directory = "page_home"
	path      = "/"
)

var html = example.Page(directory + "/page.html")

type model struct {
}

func page(c example.Context, w http.ResponseWriter, r *http.Request) (*template.Template, model, error) {
	return html, model{}, nil
}

func Get() runtime.Configuration[*runtime_chi.Runtime[example.Context]] {
	return runtime_chi.WithPage(path, http.MethodGet, page, page_error_500.Error)
}

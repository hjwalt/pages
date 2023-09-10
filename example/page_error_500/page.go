package page_error_500

import (
	"html/template"
	"net/http"

	"github.com/hjwalt/routes/example"
)

const (
	directory = "page_error_500"
)

var Html = example.Page(directory + "/page.html")

func Error(c example.Context, w http.ResponseWriter, r *http.Request, err error) *template.Template {
	return Html
}

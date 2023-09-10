package page_error_500

import (
	"html/template"
	"net/http"

	"github.com/hjwalt/routes/example"
)

const (
	pageDirectory = "page_error_500"
)

var pageTemplate = example.Page(pageDirectory + "/page.html")

func Error(c example.Context, w http.ResponseWriter, r *http.Request, err error) *template.Template {
	return pageTemplate
}

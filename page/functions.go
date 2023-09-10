package page

import (
	"context"
	"html/template"
	"net/http"
)

type Handler[C context.Context, M any] func(c C, w http.ResponseWriter, r *http.Request) (*template.Template, M, error)

type Error[C context.Context] func(c C, w http.ResponseWriter, r *http.Request, err error) *template.Template

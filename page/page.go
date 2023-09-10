package page

import (
	"context"
	"html/template"
	"net/http"
)

type Handler[C context.Context, M any] func(c C, w http.ResponseWriter, r *http.Request) (*template.Template, M, error)

type Decorator[C context.Context] func(c C, w http.ResponseWriter, r *http.Request) (C, error)

type Page[C context.Context] struct {
	Decorators []Decorator[C]
}

func (p *Page[C]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

package page

import (
	"bytes"
	"context"
	"html/template"
	"net/http"
)

func NewComponentBasic[C context.Context, M any](t *template.Template, m M) Component[C] {
	return ComponentBasic[C, M]{
		Template: t,
		Model:    m,
	}
}

type ComponentBasicModel[M any] struct {
	Model M
}

type ComponentBasic[C context.Context, M any] struct {
	Template *template.Template
	Model    M
}

func (c ComponentBasic[C, M]) Render(ctx C, w http.ResponseWriter, r *http.Request) (template.HTML, error) {
	return ComponentRender[C, M](ctx, w, r, c.Template, c.Model)
}

func ComponentRender[C context.Context, M any](ctx C, w http.ResponseWriter, r *http.Request, t *template.Template, m M) (template.HTML, error) {
	buf := new(bytes.Buffer)
	templateErr := t.Execute(buf, m)

	if templateErr != nil {
		return template.HTML(""), templateErr
	}

	return template.HTML(buf.String()), nil
}

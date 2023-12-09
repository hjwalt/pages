package page

import (
	"context"
	"html/template"
	"net/http"
)

func NewComponentMap[C context.Context, M any](t *template.Template, m M, cs map[string]Component[C]) Component[C] {
	return ComponentMap[C, M]{
		Template:   t,
		Components: cs,
		Model:      m,
	}
}

type ComponentMapModel[M any] struct {
	Components map[string]template.HTML
	Model      M
}

type ComponentMap[C context.Context, M any] struct {
	Template   *template.Template
	Components map[string]Component[C]
	Model      M
}

func (c ComponentMap[C, M]) Render(ctx C, w http.ResponseWriter, r *http.Request) (template.HTML, error) {
	return ComponentRenderMap[C, M](ctx, w, r, c.Template, c.Model, c.Components)
}

func ComponentRenderMap[C context.Context, M any](ctx C, w http.ResponseWriter, r *http.Request, t *template.Template, m M, cs map[string]Component[C]) (template.HTML, error) {
	components := map[string]template.HTML{}
	for i, com := range cs {
		var renderErr error
		components[i], renderErr = com.Render(ctx, w, r)
		if renderErr != nil {
			return template.HTML(""), renderErr
		}
	}

	model := ComponentMapModel[M]{
		Model:      m,
		Components: components,
	}

	return ComponentRender[C, ComponentMapModel[M]](ctx, w, r, t, model)
}

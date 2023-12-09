package page

import (
	"context"
	"html/template"
	"net/http"
)

func NewComponentSlice[C context.Context, M any](t *template.Template, m M, cs []Component[C]) Component[C] {
	return ComponentSlice[C, M]{
		Template:   t,
		Components: cs,
		Model:      m,
	}
}

type ComponentSliceModel[M any] struct {
	Components []template.HTML
	Model      M
}

type ComponentSlice[C context.Context, M any] struct {
	Template   *template.Template
	Components []Component[C]
	Model      M
}

func (c ComponentSlice[C, M]) Render(ctx C, w http.ResponseWriter, r *http.Request) (template.HTML, error) {
	return ComponentRenderSlice[C, M](ctx, w, r, c.Template, c.Model, c.Components)
}

func ComponentRenderSlice[C context.Context, M any](ctx C, w http.ResponseWriter, r *http.Request, t *template.Template, m M, cs []Component[C]) (template.HTML, error) {
	components := make([]template.HTML, len(cs))
	for i, com := range cs {
		var renderErr error
		components[i], renderErr = com.Render(ctx, w, r)
		if renderErr != nil {
			return template.HTML(""), renderErr
		}
	}

	model := ComponentSliceModel[M]{
		Model:      m,
		Components: components,
	}

	return ComponentRender[C, ComponentSliceModel[M]](ctx, w, r, t, model)
}

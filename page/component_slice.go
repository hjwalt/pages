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
	model, err := ComponentSliceCreateModel(ctx, w, r, c.Model, c.Components)
	if err != nil {
		return template.HTML(""), err
	}
	return ComponentRender[C, ComponentSliceModel[M]](ctx, w, r, c.Template, model)
}

func (c ComponentSlice[C, M]) Page(ctx C, w http.ResponseWriter, r *http.Request) (*template.Template, ComponentSliceModel[M], error) {
	model, err := ComponentSliceCreateModel(ctx, w, r, c.Model, c.Components)
	return c.Template, model, err
}

func ComponentSliceCreateModel[C context.Context, M any](ctx C, w http.ResponseWriter, r *http.Request, m M, cs []Component[C]) (ComponentSliceModel[M], error) {
	components := make([]template.HTML, len(cs))
	for i, com := range cs {
		var renderErr error
		components[i], renderErr = com.Render(ctx, w, r)
		if renderErr != nil {
			return ComponentSliceModel[M]{}, renderErr
		}
	}

	return ComponentSliceModel[M]{
		Model:      m,
		Components: components,
	}, nil
}

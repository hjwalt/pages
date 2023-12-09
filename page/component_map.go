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
	model, err := ComponentMapCreateModel(ctx, w, r, c.Model, c.Components)
	if err != nil {
		return template.HTML(""), err
	}
	return ComponentRender[C, ComponentMapModel[M]](ctx, w, r, c.Template, model)
}

func (c ComponentMap[C, M]) Page(ctx C, w http.ResponseWriter, r *http.Request) (*template.Template, ComponentMapModel[M], error) {
	model, err := ComponentMapCreateModel(ctx, w, r, c.Model, c.Components)
	return c.Template, model, err
}

func ComponentMapCreateModel[C context.Context, M any](ctx C, w http.ResponseWriter, r *http.Request, m M, cs map[string]Component[C]) (ComponentMapModel[M], error) {
	components := map[string]template.HTML{}
	for i, com := range cs {
		var renderErr error
		components[i], renderErr = com.Render(ctx, w, r)
		if renderErr != nil {
			return ComponentMapModel[M]{}, renderErr
		}
	}

	return ComponentMapModel[M]{
		Model:      m,
		Components: components,
	}, nil
}

package page

import (
	"context"
	"net/http"

	"github.com/hjwalt/routes/route"
	"github.com/hjwalt/runway/logger"
	"github.com/hjwalt/runway/reflect"
)

type Page[C context.Context, M any] struct {
	Decorators   []route.Decorator[C]
	PageHandler  Handler[C, M]
	ErrorHandler Error[C]
}

func (p *Page[C, M]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := reflect.Construct[C]()

	var ctxErr error
	for _, decorator := range p.Decorators {
		ctx, ctxErr = decorator(ctx, w, r)
		if ctxErr != nil {
			handleError(ctx, w, r, ctxErr, p.ErrorHandler)
			return
		}
	}

	pageTemplate, pageModel, pageErr := p.PageHandler(ctx, w, r)
	if pageErr != nil {
		handleError(ctx, w, r, pageErr, p.ErrorHandler)
		return
	}

	executeErr := pageTemplate.Execute(w, pageModel)
	if executeErr != nil {
		handleError(ctx, w, r, executeErr, p.ErrorHandler)
		return
	}
}

func handleError[C context.Context](ctx C, w http.ResponseWriter, r *http.Request, err error, errHandler Error[C]) {
	errTemplate := errHandler(ctx, w, r, err)
	errTemplateErr := errTemplate.Execute(w, err)
	if errTemplateErr != nil {
		logger.ErrorErr("error handling error", errTemplateErr)
	}
}

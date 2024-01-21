package mvc

import (
	"context"
	"net/http"

	"github.com/hjwalt/routes/route"
	"github.com/hjwalt/runway/logger"
	"github.com/hjwalt/runway/reflect"
)

func Add[C context.Context](c *route.Configuration[C], r string, m string, p Controller[C], e Error[C]) {
	routeHandler := &controller[C]{
		Decorators:   c.Decorators,
		Controller:   p,
		ErrorHandler: e,
	}
	c.AddRoute(r, m, routeHandler)
}

type Controller[C context.Context] func(c C, w http.ResponseWriter, r *http.Request) (View[C], error)

type controller[C context.Context] struct {
	Decorators   []route.Decorator[C]
	Controller   Controller[C]
	ErrorHandler Error[C]
}

func (p *controller[C]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := reflect.Construct[C]()

	var ctxErr error
	for _, decorator := range p.Decorators {
		ctx, ctxErr = decorator(ctx, w, r)
		if ctxErr != nil {
			logger.ErrorErr("error decorating", ctxErr)
			handleError(ctx, w, r, ctxErr, p.ErrorHandler)
			return
		}
	}

	view, viewErr := p.Controller(ctx, w, r)
	if viewErr != nil {
		logger.ErrorErr("error getting view", viewErr)
		handleError(ctx, w, r, viewErr, p.ErrorHandler)
		return
	}

	executeErr := view.Write(ctx, w, r)
	if executeErr != nil {
		logger.ErrorErr("error executing view", executeErr)
		handleError(ctx, w, r, executeErr, p.ErrorHandler)
		return
	}
}

package runtime_chi

import (
	"context"
	"net/http"
	"strings"

	"github.com/hjwalt/routes/page"
)

func NewConfig[C context.Context]() *HandlerConfig[C] {
	return &HandlerConfig[C]{
		routes:     map[string]*HandlerConfig[C]{},
		decorators: []page.Decorator[C]{},
	}
}

type HandlerConfig[C context.Context] struct {
	delete     http.Handler
	get        http.Handler
	post       http.Handler
	put        http.Handler
	routes     map[string]*HandlerConfig[C]
	decorators []page.Decorator[C]
}

func (config *HandlerConfig[C]) Set(method string, handler http.Handler) {
	switch method {
	case http.MethodDelete:
		config.delete = handler
	case http.MethodGet:
		config.get = handler
	case http.MethodPost:
		config.post = handler
	case http.MethodPut:
		config.put = handler
	default:
		config.get = handler
	}
}

func (config *HandlerConfig[C]) Route(fullPath string, method string, handler http.Handler) {
	path := strings.TrimSpace(fullPath)
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")

	parts := strings.Split(path, "/")

	config.Parts(parts, method, handler)
}

func (config *HandlerConfig[C]) Parts(parts []string, method string, handler http.Handler) {
	if len(parts) == 0 {
		config.Set(method, handler)
	} else {
		if curr, currExists := config.routes[parts[0]]; currExists {
			curr.Parts(parts[1:], method, handler)
		} else {
			curr := NewConfig[C]()
			config.routes[parts[0]] = curr
			curr.Parts(parts[1:], method, handler)
		}
	}
}

func (config *HandlerConfig[C]) Decorator(decorator page.Decorator[C]) {
	config.decorators = append(config.decorators, decorator)
}

func AddPage[C context.Context, M any](c *HandlerConfig[C], r string, m string, p page.Handler[C, M]) {
	routeHandler := &page.Page[C]{
		Decorators: c.decorators,
	}
	c.Route(r, m, routeHandler)
}

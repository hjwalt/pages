package route

import (
	"context"
	"net/http"
)

type Handler[C context.Context] func(c C, w http.ResponseWriter, r *http.Request)

type Decorator[C context.Context] func(c C, w http.ResponseWriter, r *http.Request) (C, error)

type Middleware func(http.Handler) http.Handler

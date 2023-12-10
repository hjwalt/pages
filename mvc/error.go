package mvc

import (
	"context"
	"net/http"

	"github.com/hjwalt/runway/logger"
)

type Error[C context.Context] func(c C, w http.ResponseWriter, r *http.Request, err error) View[C]

func handleError[C context.Context](ctx C, w http.ResponseWriter, r *http.Request, err error, errHandler Error[C]) {
	errView := errHandler(ctx, w, r, err)
	errTemplateErr := errView.Write(ctx, w, r)
	if errTemplateErr != nil {
		logger.ErrorErr("error handling error", errTemplateErr)
	}
}

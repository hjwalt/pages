package page

import (
	"context"
	"html/template"
	"net/http"

	"github.com/hjwalt/runway/logger"
)

type Error[C context.Context] func(c C, w http.ResponseWriter, r *http.Request, err error) *template.Template

func handleError[C context.Context](ctx C, w http.ResponseWriter, r *http.Request, err error, errHandler Error[C]) {
	errTemplate := errHandler(ctx, w, r, err)
	errTemplateErr := errTemplate.Execute(w, err)
	if errTemplateErr != nil {
		logger.ErrorErr("error handling error", errTemplateErr)
	}
}

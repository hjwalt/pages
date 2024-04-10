package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/hjwalt/routes/example"
	"github.com/hjwalt/routes/example/page_billing"
	"github.com/hjwalt/routes/example/page_home"
	"github.com/hjwalt/routes/runtime_chi"
	"github.com/hjwalt/runway/inverse"
	"github.com/hjwalt/runway/logger"
	"github.com/hjwalt/runway/managed"
	"github.com/hjwalt/runway/runtime"
	"github.com/joho/godotenv"
)

func main() {
	logger.Default()
	godotenv.Load()

	ctx := context.Background()
	ic := inverse.NewContainer()

	managed.AddHealth(ic)
	managed.AddHttp(ic)
	managed.AddHttpConfig(ic, map[string]string{
		managed.ConfHttpPort: "3001",
	})

	runtime_chi.AddHttpHandler[example.Context](ic)
	runtime_chi.AddMiddleware[example.Context](
		ic,
		middleware.RequestID,
		middleware.RealIP,
		middleware.CleanPath,
		middleware.Recoverer,
	)
	runtime_chi.AddRoute[example.Context](
		ic,
		runtime_chi.AddStatic[example.Context]("/static/", "./example/static"),
	)
	page_home.Add(ic)
	page_billing.Add(ic)

	managed, err := managed.New(ic, ctx)
	if err != nil {
		panic(err)
	}

	startErr := runtime.Start(
		[]runtime.Runtime{
			managed,
		},
		time.Second,
	)

	if startErr != nil {
		panic(startErr)
	}

	slog.Info("started")

	runtime.Wait()

	slog.Info("stopped")
}

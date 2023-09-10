package main

import (
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/hjwalt/routes/example"
	"github.com/hjwalt/routes/example/page_billing"
	"github.com/hjwalt/routes/example/page_home"
	"github.com/hjwalt/routes/runtime_chi"
	"github.com/hjwalt/runway/logger"
	"github.com/hjwalt/runway/runtime"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	httpRuntime := runtime_chi.New[example.Context](
		runtime_chi.WithMiddleware[example.Context](middleware.RequestID),
		runtime_chi.WithMiddleware[example.Context](middleware.RealIP),
		runtime_chi.WithMiddleware[example.Context](middleware.CleanPath),
		runtime_chi.WithMiddleware[example.Context](middleware.Recoverer),
		runtime_chi.WithPort[example.Context](3001),
		runtime_chi.WithStatic[example.Context]("/static/", "./example/static"),
		page_home.Get(),
		page_billing.Get(),
	)

	err := runtime.Start(
		[]runtime.Runtime{
			httpRuntime,
		},
		time.Second,
	)

	if err != nil {
		panic(err)
	}

	logger.Info("started")

	runtime.Wait()

	logger.Info("stopped")
}

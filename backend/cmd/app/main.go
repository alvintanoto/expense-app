package main

import (
	"context"
	"expense_app/internal/controller"
	"expense_app/internal/di"
	"expense_app/internal/util"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	if err := di.Container.Invoke(func(util util.Holder, c controller.Holder) {
		util.Logger.Info("application started")
		var (
			sig    = make(chan os.Signal, 1)
			app    = echo.New()
			parent = context.Background()
		)

		c.Routes(app)

		if err := app.Start(util.Config.Port); err != nil {
			util.Logger.Error("failed to start server")
		}

		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		util.Logger.Info("shutdown app and closing resources...")
		ctx, cancel := context.WithTimeout(parent, 30*time.Second)
		_ = app.Shutdown(ctx)
		cancel()
		util.Logger.Info("application terminated")
	}); err != nil {
		panic(err)
	}
}

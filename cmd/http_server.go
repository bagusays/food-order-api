package cmd

import (
	"context"
	"food-order-api/internal/delivery"
	"food-order-api/internal/delivery/rest"
	middl "food-order-api/internal/delivery/rest/middleware"

	"github.com/labstack/echo/v4/middleware"

	"net/http"
	"os"
	"os/signal"

	"github.com/spf13/cobra"

	"github.com/labstack/echo/v4"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	RunE:  startServer,
}

func startServer(_ *cobra.Command, _ []string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	e := echo.New()
	e.HTTPErrorHandler = middl.HTTPErrorHandler
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlAllowOrigin,
			"token",
			"Pv",
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderContentLength,
			echo.HeaderAcceptEncoding,
			echo.HeaderXCSRFToken,
			echo.HeaderXRequestID,
		},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	e.Use(middleware.Recover())

	container := delivery.NewContainer(ctx, e)
	rest.RegisterRoute(container)

	go func() {
		if err := e.Start(":8000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	return nil
}

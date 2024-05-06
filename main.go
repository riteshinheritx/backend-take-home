package main

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	flighthandler "moneda/evaluation/pkg/flight-handler"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
)

//go:embed flight_data.json
var flightDataStore []byte

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	httpPort := os.Getenv("PORT")
	e := echo.New()

	// Middleware for basic authentication
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Check if the provided username and password are correct
		if username == os.Getenv("DEFAULT_AUTH_USERNAME") && password == os.Getenv("DEFAULT_AUTH_PASSWORD") {
			return true, nil
		}
		return false, nil
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Hello, World! Please read documentation for other APIs",
		})
	})

	flight := flighthandler.NewFlight(flightDataStore)

	e.POST("/getFlightData", flight.List)

	e.POST("/searchFlightInfo", flight.Info)

	e.Any("/*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Wrong endpoint. Please read documentation",
		})
	})

	server := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}

	go func() {
		if err := e.StartH2CServer(fmt.Sprintf(":%s", httpPort), server); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal(fmt.Sprintf("Shutting down the server: %s", err.Error()))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

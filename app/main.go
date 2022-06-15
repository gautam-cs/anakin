package main

import (
	"fmt"
	"github.com/labstack/echo/v4"

	"math/rand"
	"os"
	"time"

	"gautam/server/app/config"

	"github.com/labstack/echo/v4/middleware"

	"github.com/phuslu/log"
)

//Echo Server
var Echo *echo.Echo

//Load server and set up params
func main() {
	time.Local = time.UTC

	log.DefaultLogger = log.Logger{
		Level:  log.DebugLevel,
		Caller: 1,
		Writer: log.IOWriter{os.Stdout},
	}

	log.Info().Msg("Loading config")

	//init config
	config.Load()

	log.Info().Msg("Loaded config")

	//server
	Echo := echo.New()

	log.Info().Msg("Server booting up")

	setupServer(Echo)

	addRoutes(Echo)

	log.Info().Msg("Routes Added")

	startListener(Echo)

	log.Info().Msg("Started Listener")

	rand.Seed(time.Now().UnixNano())
}

// SetupEchoServer setup echo
func setupServer(e *echo.Echo) {
	// Middleware
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisablePrintStack: config.IsDebugEnv(),
	}))

	CORSConfig := middleware.DefaultCORSConfig
	CORSConfig.AllowHeaders = []string{
		echo.HeaderCookie,
		echo.HeaderAccept,
		echo.HeaderAuthorization,
		echo.HeaderContentType,
		echo.HeaderOrigin,
		echo.HeaderXRequestedWith,
	}

	CORSConfig.AllowCredentials = true

	e.Use(middleware.CORSWithConfig(CORSConfig))
	e.Use(middleware.BodyLimitWithConfig(middleware.BodyLimitConfig{
		Limit: "1M",
	}))

	logConfig := middleware.DefaultLoggerConfig
	logConfig.Skipper = func(r echo.Context) bool {
		return r.Path() == "/status"
	}

	e.Use(middleware.LoggerWithConfig(logConfig))

	e.Use(middleware.RequestID())
	e.Use(middleware.Gzip())

	e.Validator = config.Validator()
	e.Renderer = config.Renderer()
}

//StartListener start accepting connection
func startListener(e *echo.Echo) {
	err := e.Start(fmt.Sprintf(":%d", config.ListenPort()))
	e.Logger.Fatal(err)
}

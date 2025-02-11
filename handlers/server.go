package handlers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// HTTPServer default interface for http server
type HTTPServer interface {
	Run(...string) error
}

type routes struct {
	router *gin.Engine
}

// NewServer - helps to create all api routes and http server
func NewServer(debug bool, logger *zap.Logger) HTTPServer {
	// setting up the application mode
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// router init
	r := routes{
		router: gin.New(),
	}

	// setting up the cors
	if debug {
		r.router.Use(cors.Default())
	}

	// zap logger
	// logger, _ := zap.NewProduction()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Logs all panic to error log
	r.router.Use(ginzap.RecoveryWithZap(logger, true))

	// serving static files
	r.router.Use(static.Serve("/", static.LocalFile("./static", false)))

	// creating main entry point group
	api := r.router.Group("/api")

	// creating websocket entry point group
	ws := r.router.Group("/ws")

	// grouping route with api prefix
	r.addRoute(api)
	r.addWSRoute(ws)

	return r
}

func (r routes) Run(addr ...string) error {
	return r.router.Run(addr...)
}

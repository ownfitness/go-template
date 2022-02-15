package router

import (
	"io"
	"time"

	"github.com/ownfitness/template-go/pkg/gcp"

	"github.com/ownfitness/template-go/controllers"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func New(debug bool, client *gcp.Firestore) *gin.Engine {
	// Sets gin in release mode if debug is not true
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// Set logger as middleware for gin
	r.Use(logger.SetLogger(
		logger.WithLogger(func(c *gin.Context, out io.Writer, latency time.Duration) zerolog.Logger {
			return zerolog.New(out).With().
				Str("path", c.Request.URL.Path).
				Dur("latency", latency).
				Logger()
		}),
	))

	// Health Check
	r.GET("/health", controllers.Health)

	// Users routing
	r.POST("/users", controllers.PostUser(client.Firestore))

	return r
}

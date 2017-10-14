package middleware

import (
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	iris "gopkg.in/kataras/iris.v8"
)

type Logger struct{}

func (r *Logger) Handler(ctx iris.Context) {
	path := ctx.Path()
	method := ctx.Method()

	startTime := time.Now()
	ctx.Next()
	endTime := time.Now()

	latency := endTime.Sub(startTime)
	status := strconv.Itoa(ctx.GetStatusCode())
	ip := ctx.RemoteAddr()

	log.WithFields(log.Fields{
		"status":  status,
		"latency": latency,
		"ip":      ip,
		"method":  method,
		"path":    path,
	}).Info("Request finished")
}

func NewLogger() func(iris.Context) {
	l := &Logger{}

	return l.Handler
}

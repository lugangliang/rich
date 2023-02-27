package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
	logger.Info("ss")
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	logger.Info("hello, go module", zap.ByteString("uri", ctx.RequestURI()))
}

func main() {

	logrus.Println("hello, logrus")
	logrus.Println(uuid.NewString())

	fasthttp.ListenAndServe(":8081", fastHTTPHandler)
	fmt.Println("hello world")

	return
}

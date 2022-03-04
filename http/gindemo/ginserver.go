package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const keyRequestId = "RequestId"

func main() {
	engine := gin.Default()
	log := logger()
	defer sync(log)

	engine.Use(
		func(context *gin.Context) {
			s := time.Now()
			context.Next()
			log.Info("Request processing",
				zap.String("Path", context.Request.URL.Path),
				zap.String("FullPath", context.FullPath()),
				zap.String("Status", http.StatusText(context.Writer.Status())),
				zap.String("elapsed", time.Now().Sub(s).String()),
			)
		},
		func(context *gin.Context) {
			if newUUID, err := uuid.NewUUID(); err == nil {
				context.Set(keyRequestId, newUUID.String())
			}
			context.Next()
		},
	)

	engine.GET("/ping", func(context *gin.Context) {
		response := gin.H{
			"message": "pong",
		}
		if requestId, exists := context.Get(keyRequestId); exists {
			response[keyRequestId] = requestId
		}

		context.JSON(200, response)
	})

	engine.GET("/hello", func(context *gin.Context) {
		context.String(200, "你好")
	})

	err := engine.Run()
	if err != nil {
		panic(err)
	}
}

func logger() *zap.Logger {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return log
}

func sync(logger *zap.Logger) {
	err := logger.Sync()
	if err != nil {
		panic(err)
	}
}

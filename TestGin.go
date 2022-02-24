package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	log, _ := zap.NewProduction()
	log.Info("Test info log")
	engine := gin.Default()

	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := engine.Run()
	if err != nil {
		return
	}
	log.Error(err.Error())
}

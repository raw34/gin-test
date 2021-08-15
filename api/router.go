package api

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var router *gin.Engine

func Execute()  {
	router = gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	pprof.Register(router)

	port := viper.GetString("app.port")

	err := router.Run(":" + port)

	if err != nil {
		panic(err)
	}
}

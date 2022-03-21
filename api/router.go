package api

import (
    "github.com/gin-contrib/pprof"
    "github.com/gin-gonic/gin"
    "github.com/raw34/gin-test/util"
    "github.com/spf13/viper"
)

var router *gin.Engine

func Execute() {
    router = gin.Default()
    logger := util.NewLogger()

    router.GET("/ping", func(c *gin.Context) {
        logger.Info("ping", "sssss", 233)
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    router.GET("/error", func(c *gin.Context) {
        logger.Error("normal error")
        c.JSON(500, gin.H{
            "error": "internal error",
        })
    })

    router.GET("/fatal", func(c *gin.Context) {
        logger.Fatal("fatal error")
        c.JSON(500, gin.H{
            "error": "fatal error",
        })
    })

    pprof.Register(router)

    port := viper.GetString("app.port")

    err := router.Run(":" + port)

    if err != nil {
        panic(err)
    }
}

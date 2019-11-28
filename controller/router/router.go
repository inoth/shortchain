package router

import (
	"shortchain/controller"
	mid "shortchain/controller/middleware"
	"shortchain/util/config"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.New()

	r.Use(mid.ExceptionHandle)

	mch := r.Group("/mch", mid.RequestVerify)
	mch.POST("/signin", controller.Registered)

	url := r.Group("chain", mid.RequestVerify)
	url.POST("/add", controller.GenerateShortChain)

	r.GET("/go/:shortid", controller.RedirectTo)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "ok"})
	})

	r.Run(config.Instance().ServerPort)
}

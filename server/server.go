package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"pls/config"
)

func Run(config config.Config) {
	r := gin.Default()

	r.POST("/register", registerClient)
	r.DELETE("/unregister/:clientId", unRegisterClient)
	r.GET("/owner", func(c *gin.Context) {})

	if err := r.Run(":" + config.Port); err != nil {
		log.Fatal(err)
	}
}

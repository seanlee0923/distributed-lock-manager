package server

import (
	"dlm/config"
	"github.com/gin-gonic/gin"
	"log"
)

func Run(config config.Config) {
	r := gin.Default()

	r.POST("/register", registerClient)
	r.POST("/lock/acquire", func(c *gin.Context) {})
	r.POST("/lock/release", func(c *gin.Context) {})

	if err := r.Run(":" + config.Port); err != nil {
		log.Fatal(err)
	}
}

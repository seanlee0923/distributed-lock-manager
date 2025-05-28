package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"pls/config"
)

func Run(config config.Config) {
	r := gin.Default()

	r.POST("/register", registerClient)
	r.POST("/unregister", unRegisterClient)
	r.POST("/leader", GenLeader)
	r.POST("/alive", KeepAlive)

	if err := r.Run(":" + config.Port); err != nil {
		log.Fatal(err)
	}

}

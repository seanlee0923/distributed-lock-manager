package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pls/mysql"
)

type KeepAliveRequest struct {
	ClientId   string `json:"client_id"`
	DeployName string `json:"deploy_name"`
}

func KeepAlive(c *gin.Context) {

	var req KeepAliveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "keep alive failed"})
		return
	}

	_, err := mysql.Db.Exec("UPDATE pls_clients SET last_alive = CURRENT_TIMESTAMP WHERE client_id = ? AND deploy_name = ?",
		req.ClientId, req.DeployName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "keep alive failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"client_id": req.ClientId, "deploy_name": req.DeployName})
}

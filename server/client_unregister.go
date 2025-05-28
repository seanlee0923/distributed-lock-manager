package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pls/mysql"
)

type DeleteRequest struct {
	DeployName string `json:"deploy_name"`
	ClientId   string `json:"client_id"`
}

// unRegisterClient 함수는 등록된 클라이언트를 삭제하는 함수입니다.
func unRegisterClient(c *gin.Context) {
	var req DeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "database error"})
		return
	}

	_, err := mysql.Db.Exec("DELETE FROM dlm_clients WHERE client_id = ? AND deploy_name = ?", req.ClientId, req.DeployName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"client_id": req.ClientId})

}

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"pls/mysql"
)

type RegisterRequest struct {
	DeployName string `json:"deploy_name"`
}

// registerClient 함수는 디플로이먼트 이름과 uuid를 발급하여
// 분산락 클라이언트를 디비에 등록합니다.
func registerClient(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	clientId := uuid.New().String()
	_, err := mysql.Db.Exec("INSERT INTO pls_clients (id, deploy_name) VALUES (?, ?)", clientId, req.DeployName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"client_id": clientId})
}

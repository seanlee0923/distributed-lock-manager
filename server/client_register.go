package server

import (
	"dlm/mysql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// registerClient 함수는 디플로이먼트 이름과 uuid를 발급하여
// 분산락 클라이언트를 디비에 등록합니다.
func registerClient(c *gin.Context) {
	deployName := c.Query("deployment")
	if deployName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "deployment name required"})
		return
	}

	clientId := uuid.New().String()
	_, err := mysql.Db.Exec("INSERT INTO dlm_clients (id, deploy_name) VALUES (?, ?)", clientId, deployName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"client_id": clientId})
}

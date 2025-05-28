package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pls/mysql"
)

// unRegisterClient 함수는 등록된 클라이언트를 삭제하는 함수입니다.
func unRegisterClient(c *gin.Context) {
	clientId := c.Param("clientId")
	if clientId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "clientId is empty"})
	}

	_, err := mysql.Db.Exec("DELETE FROM dlm_clients WHERE client_id = ?", clientId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"client_id": clientId})

}

package server

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pls/mysql"
)

type LeaderRequest struct {
	DeployName string `json:"deploy_name"`
	ClientId   string `json:"client_id"`
}

// GenLeader 함수는 요청한 디플로이먼트 중 누가 리더인지 선택 후
// 요청자에게 클라이언트 아이디를 응답합니다.
// 가장 최근에 등록된 파드를 리더로 선택합니다.
func GenLeader(c *gin.Context) {

	var req LeaderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "deployment name required"})
		return
	}

	var leaderId string

	leaderRow := mysql.Db.QueryRow(`SELECT id FROM pls_clients WHERE deployment_name = ? ORDER BY created_at ASC LIMIT 1`, req.DeployName)
	if err := leaderRow.Scan(&leaderId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "deployment not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db err"})
		return
	}
	log.Println(leaderId)

	c.JSON(http.StatusOK, gin.H{"client_id": leaderId})

}

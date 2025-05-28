package server

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"pls/mysql"
)

// GenLeader 함수는 요청한 디플로이먼트 중 누가 리더인지 선택 후
// 요청자에게 클라이언트 아이디를 응답합니다.
// 가장 최근에 등록된 파드를 리더로 선택합니다.
func GenLeader(c *gin.Context) {

	deployName := c.Query("deployment")
	if deployName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "deployment name required"})
		return
	}

	var leaderId string

	leaderRow := mysql.Db.QueryRow(`SELECT id FROM pls_clients WHERE deploy_name = ? ORDER BY created_at DESC LIMIT 1`, deployName)
	if err := leaderRow.Scan(&leaderId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "deployment not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db err"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"leader_id": leaderId})

}

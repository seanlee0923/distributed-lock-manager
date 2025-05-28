package server

import (
	"pls/mysql"
	"time"
)

func ZombieCollector() {
	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		_, err := mysql.Db.Exec("DELETE FROM pls_clients WHERE last_alive < (CURRENT_TIMESTAMP - INTERVAL 30 SECOND)")
		if err != nil {
			continue
		}

	}
}

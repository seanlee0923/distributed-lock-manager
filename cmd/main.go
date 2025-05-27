package main

import (
	"dlm/config"
	"dlm/mysql"
	"dlm/server"
)

func main() {

	// 환경변수 로딩해서 변수에 저장
	config.InitServerConfig()

	// mysql 연결
	mysql.Connect(config.Cfg)

	// 웹 서버 실행
	server.Run(config.Cfg)

}

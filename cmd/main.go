package main

import (
	"pls/config"
	"pls/mysql"
	"pls/server"
)

func main() {

	// 환경변수 로딩해서 변수에 저장
	config.InitServerConfig()

	// mysql 연결
	mysql.Connect(config.Cfg)

	// 웹 서버 실행
	go server.Run(config.Cfg)

	// 좀비 커넥션들 지우기
	go server.ZombieCollector()

}

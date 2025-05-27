package server

import "os"

type Config struct {
	Port          string `json:"port"`
	MysqlHost     string `json:"mysql_host"`
	MysqlPort     string `json:"mysql_port"`
	MysqlUser     string `json:"mysql_user"`
	MysqlPassword string `json:"mysql_password"`
	MysqlDatabase string `json:"mysql_database"`
}

var config Config

func InitServerConfig() {
	port := os.Getenv("DLM_PORT")
	if port == "" {
		port = "9923"
	}

	MysqlHost := os.Getenv("MYSQL_HOST")
	MysqlPort := os.Getenv("MYSQL_PORT")
	MysqlUser := os.Getenv("MYSQL_USER")
	MysqlPassword := os.Getenv("MYSQL_PASSWORD")
	MysqlDatabase := os.Getenv("MYSQL_DATABASE")
	if MysqlHost == "" || MysqlPort == "" || MysqlUser == "" || MysqlPassword == "" || MysqlDatabase == "" {
		os.Exit(0)
	}

	config.Port = port
	config.MysqlHost = MysqlHost
	config.MysqlPort = MysqlPort
	config.MysqlUser = MysqlUser
	config.MysqlPassword = MysqlPassword
	config.MysqlDatabase = MysqlDatabase
}

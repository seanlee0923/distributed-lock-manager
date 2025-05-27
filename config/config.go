package config

import (
	"log"
	"os"
)

type Config struct {
	Port          string `json:"port"`
	MysqlHost     string `json:"mysql_host"`
	MysqlPort     string `json:"mysql_port"`
	MysqlUser     string `json:"mysql_user"`
	MysqlPassword string `json:"mysql_password"`
	MysqlDatabase string `json:"mysql_database"`
}

var Cfg Config

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
		log.Fatal("failed to load config")
	}

	Cfg.Port = port
	Cfg.MysqlHost = MysqlHost
	Cfg.MysqlPort = MysqlPort
	Cfg.MysqlUser = MysqlUser
	Cfg.MysqlPassword = MysqlPassword
	Cfg.MysqlDatabase = MysqlDatabase
}

package config

import "os"

var (
	DataBaseName     = os.Getenv("MYSQL_DATABASE")
	DataBaseUserName = os.Getenv("MYSQL_USER")
	DataBasePassword = os.Getenv("MYSQL_PASSWORD")
	DataBaseHost     = os.Getenv("MYSQL_HOST")
	DataBasePort     = os.Getenv("MYSQL_PORT")
)

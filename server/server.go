package server

import (
	"gin-gorm-boilerplate/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}

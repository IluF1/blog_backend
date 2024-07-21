package main

import (
	"guthub.com/server/internal/server"
	"guthub.com/server/pkg/logger"
)

func main() {
	logger.Init()
	defer logger.Sync()
	api := server.New(":8080")
	api.Run()

}

package main

import "guthub.com/server/internal/server"

func main() {
	api := server.New(":8080")
	api.Run()
}

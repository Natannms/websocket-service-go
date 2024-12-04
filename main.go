package main

import (
	"websocket-service/cmd/web"
	"websocket-service/internal/db"
)

func main() {
	db.InitDB()
	web.StartServer()
}

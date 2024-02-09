package main

import (
	"github.com/COPPED/item-service/internal/app"
	"github.com/COPPED/item-service/internal/database"
)

func main() {
	_ = database.GetConnection()
	app.InitServer()
}

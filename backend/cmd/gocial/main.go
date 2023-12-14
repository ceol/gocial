package main

import (
	"github.com/ceol/gocial/internal/database"
)

func main() {
	database.Connect("gocial.db")
	database.Migrate()

	// server.Start()

}

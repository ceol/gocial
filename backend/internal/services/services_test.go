package services

import (
	"os"
	"testing"

	"github.com/ceol/gocial/internal/database"
)

func TestMain(m *testing.M) {
	dbSetup()
	code := m.Run()
	dbTeardown()
	os.Exit(code)
}

func dbSetup() {
	database.Connect("../../gocial_test.db")
	database.Migrate()
}

func dbTeardown() {
	database.DropTables()
	database.Disconnect()
}

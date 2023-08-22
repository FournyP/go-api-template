package settings

import (
	"log"
	"os"
	"strings"
)

type DatabaseSettings struct {
	Driver           string
	ConnectionString string
}

func NewDatabaseSettings() *DatabaseSettings {
	connectionString := os.Getenv("DATABASE_CONNECTION_STRING")

	if len(strings.TrimSpace(connectionString)) == 0 {
		log.Fatal("Please, set your DATABASE_CONNECTION_STRING env variable")
	}

	driver := os.Getenv("DATABASE_DRIVER")

	if len(strings.TrimSpace(driver)) == 0 {
		log.Fatal("Please, set your DATABASE_DRIVER env variable")
	}

	return &DatabaseSettings{
		Driver:           driver,
		ConnectionString: connectionString,
	}
}

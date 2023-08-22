package settings

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type AppSettings struct {
	Port int
}

func NewAppSettings() *AppSettings {
	portStr := os.Getenv("PORT")

	if len(strings.TrimSpace(portStr)) == 0 {
		log.Println("PORT environment variable not set. Using default port 3000.")
		portStr = "3000"
	}

	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Println("Error converting PORT environment variable to integer. Using default port 3000.")
		portStr = "3000"
	}

	return &AppSettings{
		Port: port,
	}
}

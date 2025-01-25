package main

import (
	"log"

	"github.com/LuisSilva7/theatre-reservation-api/config"
	"github.com/LuisSilva7/theatre-reservation-api/routes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := config.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// router := routes.SetupRouter(db)
	// router.Run(":8888")
}

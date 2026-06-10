package main

import (
	"context"
	"fmt"
	"log"

	"github.com/findardi/Wadi/server/internal/platform/config"
	"github.com/findardi/Wadi/server/internal/platform/database"
)

func main() {
	if err := config.LoadEnvFile("configs/.env"); err != nil {
		log.Fatal(err)
	}

	dbCfg, err := config.LoadDatabaseConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(context.Background(), dbCfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Print("Koneksi aman")
}

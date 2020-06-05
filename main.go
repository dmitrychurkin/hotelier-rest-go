package main

import (
	"log"

	"github.com/dmitrychurkin/hotelier-golang/connections"
	"github.com/dmitrychurkin/hotelier-golang/server"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	connections := connections.Init()
	defer connections.CACHE.Close()
	defer connections.DB.Close()

	err := server.Create(connections).Run()
	if err != nil {
		log.Fatal(err)
	}
}

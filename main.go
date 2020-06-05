package main

import (
	"log"

	"github.com/dmitrychurkin/hotelier-golang/connection"
	"github.com/dmitrychurkin/hotelier-golang/server"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	conn := new(connection.Connect)
	defer conn.Close()
	conn.Init()

	err := server.Create(conn).Run()
	if err != nil {
		log.Fatal(err)
	}
}

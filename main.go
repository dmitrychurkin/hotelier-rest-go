package main

import (
	"github.com/dmitrychurkin/hotelier-golang/connections"
	"github.com/dmitrychurkin/hotelier-golang/migrations"
	"github.com/dmitrychurkin/hotelier-golang/server"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	redis, db := connections.Init()
	defer redis.Close()
	defer db.Close()

	migrations.RunMigrations(db)
	server.Create().Run()
}

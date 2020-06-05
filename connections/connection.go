package connections

import (
	"fmt"
	"log"

	"github.com/dmitrychurkin/hotelier-golang/util"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Connect ...
type Connect struct {
	DB    *gorm.DB
	CACHE *redis.Client
}

// Init ...
func Init() *Connect {
	client := redis.NewClient(&redis.Options{
		Addr:     util.GetEnv("REDIS_URL"),
		Password: util.GetEnv("REDIS_PASSWORD"),
		DB:       0,
	})

	if _, err := client.Ping().Result(); err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(util.GetEnv("DB_DRIVER"), fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", util.GetEnv("DB_HOST"), util.GetEnv("DB_PORT"), util.GetEnv("DB_USER"), util.GetEnv("DB_PASSWORD"), util.GetEnv("DB_NAME"), util.GetEnv("DB_SSL_MODE")))

	if err != nil {
		log.Fatal(err)
	}

	return &Connect{db, client}
}

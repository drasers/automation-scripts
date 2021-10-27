package main

import (
	"log"
	"urlShorter/config"
	"urlShorter/handler"
	"urlShorter/storage/redis"

	"github.com/valyala/fasthttp"
)

func main() {
	configuration, err := config.FromFile("./configuration.json")
	if err != nil {
		log.Fatal(err)
	}

	service, err := redis.New(configuration.Redis.Host, configuration.Redis.Port, configuration.Redis.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer service.Close()

	router := handler.New(configuration.Options.Schema, configuration.Options.Prefix, service)

	log.Fatal(fasthttp.ListenAndServe(":"+configuration.Server.Port, router.Handler))
}

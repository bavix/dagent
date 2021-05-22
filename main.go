package main

import (
	"github.com/bavix/dagent/src/actions"
	"github.com/bavix/dagent/src/library"
	"github.com/bavix/dagent/src/store"
	"github.com/go-redis/redis/v8"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"time"
)

func main() {
	e := echo.New()
	e.Validator = library.NewValidator()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	c := jaegertracing.New(e, nil)
	defer c.Close()

	// create redis
	redisInstance := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	duration, err := time.ParseDuration(os.Getenv("REDIS_DURATION"))
	if err != nil {
		log.Fatal(err)
	}

	// make store
	storeObject := store.New(
		redisInstance,
		os.Getenv("AGENT_NAME"),
		duration)

	actMetric := actions.Metrics{Store: &storeObject}

	e.GET("/metrics", actMetric.Get)
	e.POST("/metrics", actMetric.Post)

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_ADDR")))
}

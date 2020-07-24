package main

import (
	"log"
	"os"
	"sync"

	config "github.com/joho/godotenv"
	"github.com/fajarhide/skeleton/router"
)

func main() {
	// read env
	err := config.Load(".env")
	if err != nil {
		log.Printf("can't load .env file")
		os.Exit(2)
	}
	cfgenv := os.Getenv("ENV")
	log.Printf("environment ENV=%s", cfgenv)

	service := router.MakeHandler()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		service.HTTPServerMain()
	}()

	go func() {
		defer wg.Done()
		service.GRPCServerMain()
	}()

	// Wait All services to end
	wg.Wait()
}

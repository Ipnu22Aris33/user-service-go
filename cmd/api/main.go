package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/Ipnu22Aris33/user-service-go/internal/config"
	"github.com/Ipnu22Aris33/user-service-go/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed load config:", err)
	}

	container := server.NewContainer()

	router := server.SetupRouter(container)

	fmt.Println("Server running on port", cfg.App.Port)

	if err := http.ListenAndServe(":" + cfg.App.Port, router); err != nil {
		log.Fatal(err)
	}
}
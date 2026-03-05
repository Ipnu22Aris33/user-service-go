package main

import (
	"log"
	"fmt"
	"net/http"

	"github.com/Ipnu22Aris33/user-service-go/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed load config", err)
	}
	r := gin.Default()

	r.SetTrustedProxies(nil)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.Run(":" + cfg.App.Port)
	fmt.Println("Server running on port", cfg.App.Port)
}

package api_gateway

import (
	"github.com/gin-gonic/gin"

	"log"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("❗ Falied to load config: %v", err)
	}

	r := gin.Default()
	err = r.Run(":" + cfg.Server.ServerPort)
	if err != nil {
		panic(err)
	}
}

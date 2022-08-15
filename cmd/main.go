package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ukurysheva/shop-api/pkg/auth"
	"github.com/ukurysheva/shop-api/pkg/config"
	"github.com/ukurysheva/shop-api/pkg/order"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	auServ := auth.RegisterRoutes(r, &c)
	order.RegisterRoutes(r, &c, auServ)

	r.Run(c.Port)
}

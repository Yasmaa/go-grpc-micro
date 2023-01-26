package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"api-gateway/pkg/order"
	"api-gateway/pkg/product"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}

package order

import (
	"github.com/gin-gonic/gin"
	"github.com/ukurysheva/shop-api/pkg/auth"
	"github.com/ukurysheva/shop-api/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, au *auth.ServiceClient) {
	a := auth.InitAuthMiddleWareConfig(au)

	svc := NewServiceClient(c)
	routes := r.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOrder)

}

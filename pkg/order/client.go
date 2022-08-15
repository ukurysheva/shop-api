package order

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ukurysheva/shop-api/pkg/config"
	"github.com/ukurysheva/shop-api/pkg/order/pb"
	"github.com/ukurysheva/shop-api/pkg/order/routes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func NewServiceClient(c *config.Config) *ServiceClient {
	serv, err := grpc.Dial(c.OrderServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Couldn't connect to order service")
	}

	return &ServiceClient{Client: pb.NewOrderServiceClient(serv)}
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}

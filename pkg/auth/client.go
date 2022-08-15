package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ukurysheva/shop-api/pkg/auth/pb"
	"github.com/ukurysheva/shop-api/pkg/auth/routes"
	"github.com/ukurysheva/shop-api/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func NewServiceClient(c *config.Config) *ServiceClient {
	serv, err := grpc.Dial(c.AuthServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect to auth service")
	}

	return &ServiceClient{Client: pb.NewAuthServiceClient(serv)}

}

// func NewServiceClient(c *config.Config) pb.AuthServiceClient {
// 	serv, err := grpc.Dial(c.AuthServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

// 	if err != nil {
// 		fmt.Println("Could not connect to auth service")
// 	}

// 	return pb.NewAuthServiceClient(serv)
// }

func (s *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, s.Client)
}

func (s *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, s.Client)
}

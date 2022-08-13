package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ukurysheva/shop-api/pkg/auth/pb"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx gin.Context, c pb.AuthServiceClient) {
	body := &RegisterRequestBody{}
if err := ctx.
}

package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ukurysheva/shop-api/pkg/auth/pb"
)

type AuthMiddleWareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleWareConfig(svc *ServiceClient) *AuthMiddleWareConfig {
	return &AuthMiddleWareConfig{svc: svc}
}

func (au *AuthMiddleWareConfig) AuthRequired(ctx *gin.Context) {
	header := ctx.Request.Header.Get("auth")

	if header == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(header, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	res, err := au.svc.Client.Validate(context.Background(), &pb.ValidateRequest{Token: token[1]})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()
}

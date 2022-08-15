package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ukurysheva/shop-api/pkg/auth/pb"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, cl pb.AuthServiceClient) {
	body := &RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println("Register call from api to server")
	res, err := cl.Register(context.Background(), &pb.RegisterRequest{Email: body.Email, Password: body.Password})

	if err != nil {
		fmt.Println("error regiser call - ", err.Error())
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	fmt.Println("resp: ", res)
	ctx.JSON(int(res.Status), &res)
}

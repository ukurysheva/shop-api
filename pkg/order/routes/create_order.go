package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ukurysheva/shop-api/pkg/order/pb"
)

type CreateOrderRequestBody struct {
	ProductId int64 `json:"productID"`
	Quantity  int64 `json:"quantity"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := &CreateOrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, ex := ctx.Get("userId")
	if !ex {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		UserId:    userId.(int64),
		ProductId: body.ProductId,
		Quantity:  body.Quantity,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

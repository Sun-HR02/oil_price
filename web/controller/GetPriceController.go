package controller

import (
	"github.com/gin-gonic/gin"
	"oil_price_show/model"
)

type Price struct {
	Prov        string `json:"Prov"`
	NinetyTwo   string `json:"NinetyTwo"`
	NinetyFive  string `json:"NinetyFive"`
	NinetyEight string `json:"NinetyEight"`
	Zero        string `json:"Zero"`
}

var Prices = func(ctx *gin.Context) {
	prices := model.GetPrices()
	ctx.JSON(200, prices)
}

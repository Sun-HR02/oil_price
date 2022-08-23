package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"oil_price_show/spider"
)

type Price struct {
	Prov        string `json:"Prov"`
	NinetyTwo   string `json:"NinetyTwo"`
	NinetyFive  string `json:"NinetyFive"`
	NinetyEight string `json:"NinetyEight"`
	Zero        string `json:"Zero"`
}

var Prices = func(ctx *gin.Context) {
	prices, err := spider.GetPrice()
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "error happen: %vb", err.Error())
	}

	ctx.JSON(200, prices)
}

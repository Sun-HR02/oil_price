package web

import (
	"github.com/gin-gonic/gin"
	"oil_price_show/web/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/prices", controller.Prices)
	return r
}

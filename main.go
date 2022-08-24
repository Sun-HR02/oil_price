package main

import (
	"log"
	"oil_price_show/conf"
	"oil_price_show/middleware"
	"oil_price_show/model"
	"oil_price_show/web"
	"time"
)

func main() {
	r := web.NewRouter()
	go func() {
		model.UpdatePrices()
		for {
			time.Sleep(24 * 60 * time.Second)
			log.Println("Oil prices updated")
			model.UpdatePrices()
		}
	}()

	r.Use(middleware.Cors())

	appPort := conf.AppPort()
	err := r.Run(":" + appPort)
	if err != nil {
		log.Println(err)
	}

}

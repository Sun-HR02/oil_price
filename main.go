package main

import (
	"log"
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
			log.Println("A day passed. Update prices.")
			model.UpdatePrices()
		}
	}()

	r.Run(":8080")

}

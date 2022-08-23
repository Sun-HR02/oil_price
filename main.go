package main

import "oil_price_show/web"

func main() {
	r := web.NewRouter()

	r.Run(":8080")
}

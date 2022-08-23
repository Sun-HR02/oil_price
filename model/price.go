package model

import (
	"log"
	"oil_price_show/spider"
)

type Price struct {
	Prov        string
	NinetyTwo   string
	NinetyFive  string
	NinetyEight string
	Zero        string
}

var conn = NewConnection()

func CreatePrices() {
	conn.AutoMigrate(&Price{})

	prices, err := spider.GetPrice()
	if err != nil {
		log.Println(err)
	}
	for _, i := range prices {
		conn.Create(Price{
			Prov:        i.Prov,
			NinetyTwo:   i.NinetyTwo,
			NinetyFive:  i.NinetyFive,
			NinetyEight: i.NinetyEight,
			Zero:        i.Zero,
		})
	}
}

func GetPrices() []Price {
	var Prices []Price

	conn.Find(&Prices)
	return Prices
}

func UpdatePrices() {
	prices, err := spider.GetPrice()
	if err != nil {
		log.Println(err)
	}
	for _, i := range prices {
		p := Price{Prov: i.Prov}
		conn.Model(&p).Where("prov = ?", i.Prov).Updates(Price{
			NinetyTwo:   i.NinetyTwo,
			NinetyFive:  i.NinetyFive,
			NinetyEight: i.NinetyEight,
			Zero:        i.Zero,
		})
	}

}

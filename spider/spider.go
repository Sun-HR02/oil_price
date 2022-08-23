package spider

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

type Price struct {
	Prov        string
	NinetyTwo   string
	NinetyFive  string
	NinetyEight string
	Zero        string
}

func GetPrice() ([]Price, error) {
	c := colly.NewCollector()
	prices := make([]Price, 0)
	// Find and visit all links
	c.OnHTML(".content", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, element *colly.HTMLElement) {
			var price Price
			price.Prov = GetProv(i)
			if price.Prov != "Not Found" {
				element.ForEach("td", func(i int, htmlElement *colly.HTMLElement) {
					if i == 1 {
						price.NinetyTwo = htmlElement.Text
					}
					if i == 2 {
						price.NinetyFive = htmlElement.Text
					}
					if i == 3 {
						price.NinetyEight = htmlElement.Text
					}
					if i == 4 {
						price.Zero = htmlElement.Text
					}
				})
			}
			if price.Prov != "Not found" {
				prices = append(prices, price)
			}

		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit("https://oil.usd-cny.com/")
	if err != nil {
		return nil, err
	}
	return prices, nil
}

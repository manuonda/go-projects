package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {

	c := colly.NewCollector(colly.AllowedDomains("j2store.net"))

	c.OnHTML("div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		fmt.Print(h.Text)
	})

	c.Visit("http://j2store.net/demo/index.php/shop")

}

package main

import (
	"fmt"
	"newsfeeder/models"
)

func main() {
	feed := models.New()
	fmt.Println(feed)

	feed.Add(models.Item{"Hello", "How are you?"})

	fmt.Println(feed)
}

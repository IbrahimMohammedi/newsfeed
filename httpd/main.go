package main

import (
	"fmt"
	"newsfeeder/platform/newsfeed"
)

func main() {
	//r := gin.Default()
	//r.GET("/ping", handler.PingGet())

	//r.Run()

	feed := newsfeed.New()

	fmt.Println(feed)

	feed.Add(newsfeed.Item{"sah", "dude"})

	fmt.Println(feed)

}

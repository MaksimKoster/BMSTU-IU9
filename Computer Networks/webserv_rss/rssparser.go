package main

import (
	"fmt"
	"github.com/RealJK/rss-parser-go"
	"log"
)

func main() {
	_ = getRss()
}

func getRss() *rss.RSS {
	rssObject, err := rss.ParseRSS("https://meduza.io/rss/all")

	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Title           : %s\n", rssObject.Channel.Title)
	fmt.Printf("Generator       : %s\n", rssObject.Channel.Generator)
	fmt.Printf("PubDate         : %s\n", rssObject.Channel.PubDate)
	fmt.Printf("LastBuildDate   : %s\n", rssObject.Channel.LastBuildDate)
	fmt.Printf("Description     : %s\n", rssObject.Channel.Description)

	fmt.Printf("Number of Items : %d\n", len(rssObject.Channel.Items))

	for v, item := range rssObject.Channel.Items {
		//item := rssObject.Channel.Items[v]
		fmt.Println()
		fmt.Printf("Item Number : %d\n", v)
		fmt.Printf("Title       : %s\n", item.Title)
		fmt.Printf("Link        : %s\n", item.Link)
		fmt.Printf("Description : %s\n", item.Description)
		fmt.Printf("Guid        : %s\n", item.Guid.Value)
	}

	return rssObject
}

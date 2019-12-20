package main

import (
	"fmt"
	"golang.org/x/net/html"
	_ "html/template"
	"net/http"

	"log"
)

func getChildren(node *html.Node) []*html.Node {
	var children []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		children = append(children, c)
	}
	return children
}

func getAttr(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func isText(node *html.Node) bool {
	return node != nil && node.Type == html.TextNode
}

func isElem(node *html.Node, tag string) bool {
	return node != nil && node.Type == html.ElementNode && node.Data == tag
}

func isDiv(node *html.Node, class string) bool {
	return isElem(node, "div") && getAttr(node, "class") == class
}

func readItem(item *html.Node) *Item {
	if a := item.FirstChild; isElem(a, "a") {
		cs := getChildren(a)
		if isText(cs[0]) {
			return &Item{
				Ref:   getAttr(a, "href"),
				Time:  getAttr(cs[0], "title"),
				Title: cs[0].Data,
			}
		}
	}
	return nil
}

type Item struct {
	Ref, Time, Title string
}

func downloadNews(link string) []*Item {
	log.Printf("sending request to %s", link)
	if response, err := http.Get("http://" + link); err != nil {
		log.Panic(err)
	} else {
		defer response.Body.Close()
		status := response.StatusCode
		log.Printf("got response from %s : %d", link, status)
		if status == http.StatusOK {
			if doc, err := html.Parse(response.Body); err != nil {
				log.Panic("invalid HTML from ", link, " error ", err)
			} else {
				log.Printf("HTML from %s parsed successfully", link)
				items := search(doc)
				log.Println(len(items))
				return items
			}
		}
	}
	return nil
}

func search(node *html.Node) []*Item {
	if isDiv(node, "b-yellow-box__wrap") {
		log.Println("====b-yellow-box__wrap check + ====")
		var items []*Item
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if isDiv(c, "item") {
				item := readItem(c)
				if item != nil {
					log.Println("appending, len = ", len(items))
					items = append(items, item)
				}
			}
		}
		return items
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if items := search(c); items != nil {
			return items
		}
	}
	return nil
}

func printMainNews(items []*Item) {
	for _, v := range items {
		fmt.Println(v.Title)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		http.ServeFile(w, r, "index.html")
	} else {
		_, _ = fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

//===================================================================================================

func main() {
	log.Println("Downloader started")
	printMainNews(downloadNews("lenta.ru"))
	//for v, _ := range items {
	//	fmt.Println(v)
	//}
}

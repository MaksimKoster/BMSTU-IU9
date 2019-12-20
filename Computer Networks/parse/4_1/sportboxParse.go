package main

import (
	"fmt"
	"golang.org/x/net/html"
	_ "html/template"
	"net/http"
	"strings"

	"log"
)

func getChildren_(node *html.Node) []*html.Node {
	var children []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if getAttr_(c, "class") == "" {
			continue
		}
		children = append(children, c)
	}
	return children
}

func getAttr_(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func isElem_(node *html.Node, tag string) bool {
	return node != nil && node.Type == html.ElementNode && node.Data == tag
}

func isDiv_(node *html.Node, class string) bool {
	return isElem_(node, "div") && getAttr_(node, "class") == class
}

func readItem_(item *html.Node) *Item_ {
	c := item.FirstChild
	attr := strings.Split(getAttr_(c, "class"), " ")
	if attr[0] == "sport-title" {
		return &Item_{
			Left:   c.FirstChild.FirstChild.Data,
			Middle: "",
			Right:  "TITLE",
		}
	}
	var gamesCh []*html.Node
	live := false
	if getAttr_(c.FirstChild.NextSibling, "class") != "games" {
		gamesCh = getChildren_(c.FirstChild.NextSibling.NextSibling.NextSibling)
		live = true
	} else {
		gamesCh = getChildren_(c.FirstChild.NextSibling)
	}
	if gamesCh == nil {
		gamesCh = getChildren_(c.FirstChild.NextSibling.FirstChild.NextSibling)
	}
	if len(gamesCh) == 3 {
		return &Item_{
			Left:   getEdge(gamesCh[0]),
			Middle: getMiddle(gamesCh[1].FirstChild, live),
			Right:  getEdge(gamesCh[2]),
		}
	} else if len(gamesCh) == 2 {
		return getF1Result(gamesCh[0])
	}
	return nil
}

func getF1Result(item *html.Node) *Item_ {
	var middle []string
	for c := item.Parent.NextSibling.NextSibling.FirstChild.NextSibling; c != nil; c = c.NextSibling {
		if c.Data != "li" {
			continue
		}
		middle = append(middle, fmt.Sprintf("%s", getTable(c)))
	}
	return &Item_{
		Left:   item.FirstChild.Data,
		Middle: strings.Join(middle, "\n"),
		Right:  "F1",
	}
}

func getTable(item *html.Node) string {
	res := ""
	for c := item.FirstChild; c != nil; c = c.NextSibling {
		if c.Attr != nil {
			switch getAttr_(c, "class") {
			case "place":
				res += c.FirstChild.Data
			case "name":
				res += fmt.Sprintf(" %s", c.FirstChild.Data)
			case "result":
				res += fmt.Sprintf(" (%s)", c.FirstChild.Data)
			}
		}
	}
	return res
}

func getEdge(item *html.Node) string {
	tmp := item.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild
	if tmp == nil {
		return ""
	}
	if getAttr_(tmp.Parent, "class") == "icons" {
		tmp = tmp.Parent.PrevSibling.PrevSibling.FirstChild
	}
	return tmp.Data
}

func getMiddle(item *html.Node, live bool) string {
	res := ""
	attr := getAttr_(item, "class")
	if attr == "time" {
		res = fmt.Sprintf("starts at %s:%s", item.FirstChild.Data, item.FirstChild.NextSibling.FirstChild.Data)
	} else if attr == "score" {
		res = item.FirstChild.Data
	} else if attr == "mdash gray" {
		res = getTennisResults(item)
	}
	if live {
		res += " [LIVE]"
	}
	return res
}

func getTennisResults(item *html.Node) string {
	return strings.TrimSpace(item.Parent.Parent.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data)
}

func search_(node *html.Node) []*Item_ {
	if isDiv_(node, "clearfix") {
		//log.Println("====clearfix check + ====")
		var items []*Item_
		for c := node.NextSibling; c != nil; c = c.NextSibling {
			if isDiv_(c, "col-lg-6 col-md-6 col-sm-6 col-xs-12") {
				item := readItem_(c)
				if item != nil {
					items = append(items, item)
				} else if len(items) > 0 {
					log.Println(items[len(items)-1])
				}
			}
		}
		return items
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if items := search_(c); items != nil {
			return items
		}
	}
	return nil
}

type Item_ struct {
	Left, Middle, Right string
}

func downloadResults(link string) []*Item_ {
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
				items := search_(doc)
				return items
			}
		}
	}
	return nil
}

//===================================================================================================

func main() {
	log.Println("Downloader started")
	res := downloadResults("sportbox.ru")
	for _, item := range res {
		if item.Right == "F1" {
			fmt.Printf("\n%s\n%s", item.Left, item.Middle)
		} else if item.Right == "TITLE" {
			fmt.Printf("\n\n%s", item.Left)
		} else {
			fmt.Printf("\n%s : %s - %s", item.Left, item.Right, item.Middle)
		}
	}
}

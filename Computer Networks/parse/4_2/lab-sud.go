package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

type Item struct {
	Ref, Time, Title string
}

func parseLabSud() int {
	if response, err := http.Get("http://lab-sud.ru"); err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		status := response.StatusCode
		if status == http.StatusOK {
			if doc, err := html.Parse(response.Body); err != nil {
				log.Panic("invalid HTML from ", "sitemap", " error ", err)
			} else {
				log.Printf("HTML from %s parsed successfully", "sitemap")
				items := searchSite(doc)
				items = append(items, searchOtherLinks(doc)...)
				log.Println(len(items))
				return len(items)
			}
		}
	}
	return 0
}

func searchOtherLinks(node *html.Node) []*Item {
	for c := node.FirstChild; c != nil; c = c.NextSibling {

	}
	return nil
}

func parseSubSite(link string) []*Item {
	if response, err := http.Get("http://" + link); err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		if response.StatusCode == http.StatusOK {
			if doc, err := html.Parse(response.Body); err != nil {
				log.Fatal("SubSite parse error")
			} else {
				return searchSubSite(doc)
			}
		}
	}
	return nil
}

func searchSubSite(node *html.Node) []*Item {
	if isElem(node, "li") && getAttr(node, "class") == "next" {
		log.Println("li next +++")
		node = node.FirstChild
		var items []*Item
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if isElem(c, "li") {
				if a := c.FirstChild; a.Data == "a" {
					item := &Item{
						Ref:   "",
						Time:  "",
						Title: fmt.Sprintf("lab-sud.ru%s", getAttr(a, "href")),
					}
					items = append(items, item)
				}
			}
		}
		return items
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if items := searchSubSite(c); items != nil {
			return items
		}
	}
	return nil
}

func searchSite(node *html.Node) []*Item {
	if isUl(node, "rubrics") {
		log.Println("col-sm-4")
		var items []*Item
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if isElem(c, "li") {
				item := readSite(c)
				if item != nil {
					log.Println("appending, len = ", len(items))
					items = append(items, item)
					sub := parseSubSite(item.Title)
					if sub != nil {
						items = append(items, sub...)
					}
				}
			}
		}
		return items
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if items := searchSite(c); items != nil {
			return items
		}
	}
	return nil
}

func main() {
	sitemapSize, labSudSize := 0, parseLabSud()
	if sitemap, err := http.Get("http://lab-sud.ru/sitemap.xml"); err != nil {
		log.Fatal(err)
	} else {
		defer sitemap.Body.Close()
		sitemapSize = parseSitemap(sitemap)
	}
	if sitemapSize == labSudSize {
		fmt.Println("OK")
	} else {
		fmt.Printf("xml size: %d, lab-sud size: %d, not OK", sitemapSize, labSudSize)
	}
}

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

func isUl(node *html.Node, class string) bool {
	return isElem(node, "ul") && getAttr(node, "class") == class
}

func readSite(item *html.Node) *Item {
	if a := item.FirstChild; isElem(a, "a") {
		if href := getAttr(a, "href"); href != "" {
			return &Item{
				Ref:   "-",
				Time:  "-",
				Title: fmt.Sprintf("lab-sud.ru%s", href),
			}
		}
	}
	return nil
}

func readUrl(item *html.Node) *Item {
	if a := item.FirstChild; isElem(a, "loc") {
		cs := getChildren(a)
		if isText(cs[0]) {
			return &Item{
				Ref:   "-",
				Time:  "-",
				Title: cs[0].Data,
			}
		}
	}
	return nil
}

func searchSitemap(node *html.Node) []*Item {
	if node.Data == "urlset" {
		log.Println("urlset")
		var items []*Item
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if c.Data == "url" {
				item := readUrl(c)
				if item != nil {
					items = append(items, item)
				}
			}
		}
		return items
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if items := searchSitemap(c); items != nil {
			return items
		}
	}
	return nil
}

func parseSitemap(sitemap *http.Response) int {
	status := sitemap.StatusCode
	if status == http.StatusOK {
		if doc, err := html.Parse(sitemap.Body); err != nil {
			log.Panic("invalid HTML from ", "sitemap", " error ", err)
		} else {
			log.Printf("HTML from %s parsed successfully", "sitemap")
			items := searchSitemap(doc)
			log.Println(len(items))
			return len(items)
		}
	}
	return -1
}
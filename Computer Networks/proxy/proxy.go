package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var addres = "localhost"

func ExtractText(node *html.Node) string {
	var content strings.Builder
	for _, text := range GetChildren(node) {
		if html.TextNode == text.Type {
			content.WriteString(text.Data)
		} else {
			children := GetChildren(text)
			if len(children) > 0 {
				content.WriteString(ExtractText(text))
			}
		}
	}
	return content.String()
}

func GetChildren(node *html.Node) []*html.Node {
	var children []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		children = append(children, c)
	}
	return children
}

func GetAttr(node *html.Node, key string) (string, error) {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val, nil
		}
	}
	return "", errors.New("no such attribute in node")
}

func GetByTag(node *html.Node, tag string) []*html.Node {
	if node != nil {
		nodes := make([]*html.Node, 0)
		if node.Data == tag {
			nodes = append(nodes, node)
		}
		for _, c := range GetChildren(node) {
			subNodes := GetByTag(c, tag)
			if subNodes != nil {
				nodes = append(nodes, subNodes...)
			}
		}
		return nodes
	}
	return nil
}

func GetByClass(node *html.Node, class string) []*html.Node {
	if node != nil {
		nodes := make([]*html.Node, 0)
		nodeClass, _ := GetAttr(node, "class")
		if nodeClass == class {
			nodes = append(nodes, node)
		}
		for _, c := range GetChildren(node) {
			subNodes := GetByClass(c, class)
			if subNodes != nil {
				nodes = append(nodes, subNodes...)
			}
		}
		return nodes
	}
	return nil
}

func GetById(node *html.Node, id string) []*html.Node {
	if node != nil {
		nodes := make([]*html.Node, 0)
		nodeId, _ := GetAttr(node, "id")
		if nodeId == id {
			nodes = append(nodes, node)
		}
		for _, c := range GetChildren(node) {
			subNodes := GetById(c, id)
			if subNodes != nil {
				nodes = append(nodes, subNodes...)
			}
		}
		return nodes
	}
	return nil
}

func GetHTML(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	utf8, err := charset.NewReader(response.Body, response.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(utf8)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetDocument(url string) (*html.Node, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	utf8, err := charset.NewReader(response.Body, response.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	node, err := html.Parse(utf8)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func getSite(r *http.Request) string {
	site, err := r.Cookie("site")
	if err != nil || !strings.Contains(r.URL.Path, site.Value) {
		split := strings.Split(r.URL.Path, "/")
		if len(split) == 1 {
			return ""
		}
		path := split[1]
		r.AddCookie(&http.Cookie{
			Name:  "site",
			Value: path,
		})
		return path
	}
	return site.Value
}

func get(w http.ResponseWriter, r *http.Request) {
	requested := r.URL.String()
	var err error
	var doc *html.Node
	if !strings.Contains(requested, "http") || !strings.Contains(requested, "https") {
		doc, err = GetDocument("https:/" + requested)
		if err != nil {
			doc, err = GetDocument("http:/" + requested)
			if err != nil {
				_, _ = w.Write([]byte(err.Error()))
				return
			}
		}
	} else {
		doc, err = GetDocument(requested)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
		}
	}

	refs := GetByTag(doc, "a")
	if err != nil {
		fmt.Println(err.Error())
	}

	site := getSite(r)
	for _, ref := range refs {
		for i, attr := range ref.Attr {
			if attr.Key == "href" {
				val := attr.Val

				if !strings.Contains(val, "://") {
					if len(strings.Split(val, "/")) > 1 && strings.Split(val, "/")[1] != site {
						if strings.Contains(val, ".js") {
							ref.Attr[i].Val = site + ref.Attr[i].Val
						} else if strings.Index(val, "/") == 0 {
							ref.Attr[i].Val = "http://" + addres + "/" + site + ref.Attr[i].Val
						}
					}
				} else {
					protocol := strings.Split(ref.Attr[i].Val, "://")[0]
					part := strings.Split(ref.Attr[i].Val, "://")[1]
					ref.Attr[i].Val = protocol + "://" + addres + "/" + part
				}
			}
		}
	}

	err = html.Render(w, doc)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	addres += ":8005"
	http.HandleFunc("/", get)
	if err := http.ListenAndServe(":8005", nil); err != nil {
		log.Fatal(err)
	}
}

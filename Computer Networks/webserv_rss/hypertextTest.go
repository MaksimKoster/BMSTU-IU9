package main

import (
	"fmt"
	"github.com/RealJK/rss-parser-go"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

type byValue rss.Channel

func main() {
	http.HandleFunc("/", htmlTest)
	err := http.ListenAndServe(":8085", nil)
	log.Println("Starting server...")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func htmlTest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		writeHtml()
		http.ServeFile(w, r, "index.html")
	} else {
		_, _ = fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func writeHtml() {
	_ = os.Remove("index.html")
	file, err := os.Create("index.html")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	//file.
	_, _ = file.Write([]byte("<!DOCTYPE > \n " +
		"<html>\n " +
		"<head> \n\t <meta charset=\"UTF-8\"> \n</head>\n" +
		"<body>"))

	rssObjects := make([]rss.RSS, 0)

	rssObjectMeduza := getRss_("https://meduza.io/rss/all")
	rssObjectSports := getRss_("https://www.sports.ru/rss/rubric.xml?s=208")
	rssObjectZonaM := getRss_("https://zona.media/rss")

	rssObjects = append(rssObjects, *rssObjectMeduza, *rssObjectZonaM, *rssObjectSports)
	rssUnion := unify(rssObjects)

	for _, item := range rssUnion.Channel.Items {
		t, err := item.PubDate.Parse()
		if err != nil {
			t, err = time.Parse(time.RFC1123Z, string(item.PubDate))
			if err != nil {
				log.Println(err)
			}
		}
		//fmt.Println(time)
		_, _ = file.Write([]byte("\n\t<a>" + item.Author + ": &nbsp;</a>\n"))
		_, _ = file.Write([]byte("\t<a href=\"" + item.Link + "\">" + item.Title + "&nbsp;</a>\n"))
		_, _ = file.Write([]byte("\t<a>" + t.Format("2006-01-02 15:04:05") + "<br></a>\n"))
	}
	_, _ = file.Write([]byte("<body> \n </html>"))
}

func getRss_(url string) *rss.RSS {
	rssObject, err := rss.ParseRSS(url)
	if err != nil {
		log.Println(err)
	}
	return rssObject
}

func unify(objs []rss.RSS) *rss.RSS {
	var result rss.RSS
	for _, obj := range objs {
		fmt.Println(obj.Channel.Title)
		for _, item := range obj.Channel.Items {
			result.Channel.Items = append(result.Channel.Items, item)
			size := len(result.Channel.Items) - 1
			result.Channel.Items[size].Author = obj.Channel.Title
		}
	}

	sort.Slice(result.Channel.Items, func(i, j int) bool {
		return result.Channel.Items[i].PubDate > result.Channel.Items[j].PubDate
	})
	//sort.Sort(byValue(result.Channel))

	return &result
}

func (f byValue) Len() int {
	return len(f.Items)
}
func (f byValue) Less(i, j int) bool {
	t1, err := f.Items[i].PubDate.Parse()
	if err != nil {
		t1, _ = time.Parse(time.RFC1123Z, string(f.PubDate))
	}
	t2, err := f.Items[j].PubDate.Parse()
	if err != nil {
		t2, _ = time.Parse(time.RFC1123Z, string(f.PubDate))
	}
	return t1.Before(t2)
}

func (f byValue) Swap(i, j int) {
	f.Items[i], f.Items[j] = f.Items[j], f.Items[i]
}

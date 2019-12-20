package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", HomeRouterHandler_)
	err := http.ListenAndServe(":9005", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func HomeRouterHandler_(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	if r.Method == "GET" {
		WriteHtmlForTest()
		http.ServeFile(w, r, "form.html")
	}
}

func WriteHtmlForTest() {
	_ = os.Remove("form.html")
	file, err := os.Create("form.html")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	_, _ = file.Write([]byte("<!DOCTYPE > \n " +
		"<html>\n " +
		"<head> \n\t <meta charset=\"UTF-8\"> \n</head>\n" +
		"<body>"))

	_, _ = file.Write([]byte("<a href=\"https://google.com\">GOOGLE<br></a>\n" +
		"<a href=\"https://yandex.ru\">YANDEX<br></a>\n" +
		"<a href=\"https://mail.ru\">MAILRU<br></a>\n"))

	_, _ = file.Write([]byte("<body> \n </html>"))
}

func SendFormHtml() {

}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	log.Println("rForm: ", r.Form)
	log.Println("path", r.URL.Path)
	log.Println("scheme", r.URL.Scheme)
	log.Println("method", r.Method)
	log.Println("[\"url_long\"]", r.Form["url_long"])
	for k, v := range r.Form {
		log.Println("key:", k)
		log.Println("value:", strings.Join(v, " "))
	}
	_, _ = fmt.Fprint(w, "Hello, World!")
}

package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

type info struct {
	login string
	pass  string
}

type tests struct {
	Test1 string
	Test2 string
	Test3 string
}

var (
	serv = "lab.posevin.com:22"
	pk   = flag.String("key", defaultKeyPath(), "Private key file")
	akks = []info{{
		login: "test1",
		pass:  "123",
	}, {
		login: "test2",
		pass:  "123",
	}, {
		login: "test3",
		pass:  "123",
	}}
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/update", updateHandler)
	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.Fatal(err)
	}
}

func exec(lIn info, cmd string, signer ssh.Signer, i int) string {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	log.Println(lIn.login, " ", lIn.pass)
	config := &ssh.ClientConfig{
		Config: ssh.Config{},
		User:   lIn.login,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
			ssh.Password(lIn.pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", serv, config)
	if err != nil {
		panic(fmt.Sprintf("err at %d: %s", i, err.Error()))
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		log.Panic(err)
	}
	defer session.Close()
	b, err := session.Output(cmd)
	if err != nil {
		log.Panic(err.Error())
	}
	if string(b) == "" {
		fmt.Printf("\"%s\": success", cmd)
	}
	fmt.Printf("%s: %d\n", lIn.login, len(string(b)))
	arr := strings.Split(string(b), "\n")
	var res = ""
	var space = regexp.MustCompile("\\s+")
	for _, a := range arr {
		strs := space.Split(a, -1)
		for i, s := range strs {
			if i == 0 || i == 4 {
				log.Println("i: ",i, " s: ",s)
				res += fmt.Sprintf("%s\t", s)
			}
		}
		res += "\n"
	}
	return res
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	var df = make(map[string]chan string)
	df["test1"] = make(chan string)
	df["test2"] = make(chan string)
	df["test3"] = make(chan string)
	key, err := ioutil.ReadFile(*pk) // получение приватного ключа для авторизации на ssh
	if err != nil {
		log.Panic(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal(err)
	}
	var cmd = "df -h"
	for i, akk := range akks {
		go func(i int, akk info) {
			df[akk.login] <- exec(akk, cmd, signer, i)
		}(i, akk)
	}
	time.Sleep(time.Second / 2)
	var data = tests{
		Test1: <-df["test1"],
		Test2: <-df["test2"],
		Test3: <-df["test3"],
	}
	tmplt := template.Must(template.ParseFiles("index.html"))
	_ = tmplt.Execute(w, data)
}

func defaultKeyPath() string {
	home := os.Getenv("HOME")
	if len(home) > 0 {
		return path.Join(home, ".ssh/id_rsa")
	}
	return ""
}

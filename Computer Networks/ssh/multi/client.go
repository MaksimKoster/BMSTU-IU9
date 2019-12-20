package main

import (
	"flag"
	"fmt"
	"github.com/skorobogatov/input"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

type info struct {
	login string
	pass  string
}

var (
	serv = "lab.posevin.com:22"
	pk   = flag.String("key", defaultKeyPath(), "Private key file")
	akks = []info{{
		login: "test1",
		pass:  "12345678990test1",
	}, {
		login: "test2",
		pass:  "12345678990test2",
	}, {
		login: "test3",
		pass:  "12345678990test3",
	}}
)

func main() {
	key, err := ioutil.ReadFile(*pk) // получение приватного ключа для авторизации на ssh
	if err != nil {
		log.Panic(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal(err)
	}
	var cmd string
	cmd = input.Gets()
	cmd += " && echo"
	for i, akk := range akks {
		log.Printf("%s %s %d", akk, cmd, i)
		go func(i int, akk info) {
			exec(akk, cmd, signer, i)
		}(i, akk)
	}
	select {
	case <-time.After(time.Second * 5):
		log.Println("Expired")
	}
}

func exec(lIn info, cmd string, signer ssh.Signer, i int) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
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
		log.Panic(fmt.Sprintf("err at %d", i))
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
	fmt.Printf("%s: %s", lIn.login, b)
}

func defaultKeyPath() string {
	home := os.Getenv("HOME")
	if len(home) > 0 {
		return path.Join(home, ".ssh/id_rsa")
	}
	return ""
}

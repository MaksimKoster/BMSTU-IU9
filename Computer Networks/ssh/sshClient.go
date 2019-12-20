package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/skorobogatov/input"
	"setiLabs/config"
)

var (
	pk       = flag.String("key", defaultKeyPath(), "Private key file")
	user     = flag.String("user", config.GetSshLogin(), "User name")
	password = flag.String("pwd", config.GetSshPassword(), "Password")
	//host      = flag.String("host", "185.20.227.83", "Host")
	//port      = flag.String("port", "22", "Port")
)

func main() {
	flag.Parse()
	key, err := ioutil.ReadFile(*pk) // получение приватного ключа для авторизации на ssh
	if err != nil {
		log.Panic(err)
	}

	signer, err := ssh.ParsePrivateKey(key) //  проверка ключа
	if err != nil {
		log.Panic(err)
	}
	config := &ssh.ClientConfig{
		Config: ssh.Config{},
		User:   *user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
			ssh.Password(*password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	fmt.Println(signer)

	hosts := os.Args[1:]

	var cmd string

	result := make(chan string, 10)
	cmd = input.Gets()
	for _, host := range hosts {
		client, err := ssh.Dial("tcp", host, config)
		if err != nil {
			log.Panic(err)
		}
		go func(host string) {
			//defer client.Close()
			//execute(cmd, host, config)
			result <- fmt.Sprintf("%s ::\n%s::\n", host, execute(cmd, host, client))
		}(host)
	}

	for range hosts {
		select {
		case res := <-result:
			fmt.Println(res)
		case <-time.After(time.Second * 20):
			fmt.Println("Timed out")
		}
	}
}

func execute(cmd, addr string, client *ssh.Client) string {
	//client, err := ssh.Dial("tcp", addr, config)
	//if err != nil {
	//	log.Panic(err)
	//}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		log.Panic(err)
	}
	defer session.Close()
	b, err := session.Output(cmd)
	if err != nil {
		return err.Error()
	}
	if string(b) == "" {
		return fmt.Sprintf("\"%s\": success", cmd)
	}
	return string(b)
	//********************************************
	//modes:=ssh.TerminalModes{
	//	ssh.ECHO:0,
	//	ssh.TTY_OP_ISPEED:14400,
	//	ssh.TTY_OP_OSPEED:14400,
	//}
	//if err:=session.RequestPty("xterm",80,40,modes);err!=nil{
	//	handleError(session.Close(),false)
	//	handleError(err,true)
	//}
	//in, err := session.StdinPipe()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer in.Close()
	//handleError(err,true)
	//
	//session.Stdout = os.Stdout
	//
	//err = session.Shell()
	//if err = session.Shell(); err != nil {
	//	log.Println(err)
	//}
	//_,err=fmt.Fprintf(in,"%s\n",cmd)
	//if err != nil {
	//	log.Println(err)
	//}
	//_ = session.Wait()
	//************************************************
}

func defaultKeyPath() string {
	home := os.Getenv("HOME")
	if len(home) > 0 {
		return path.Join(home, ".ssh/id_rsa")
	}
	return ""
}

func handleError(err error, priority bool) bool {
	if err != nil {
		if priority {
			panic(err)
		} else {
			log.Fatal(err)
			return false
		}
	}
	return true
}

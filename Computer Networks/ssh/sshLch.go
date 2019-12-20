package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os"
	config2 "setiLabs/config"
)

func handleError_(err error, priority bool) bool {
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

func main() {

	var (
		user, password string
	)

	flag.StringVar(&user, "user", "iu9_32_05", "user name")
	flag.StringVar(&password, "password", config2.GetSshPassword(), "password")
	//flag.StringVar(&host,"host","localhost","host name")
	//flag.IntVar(&port,"port",2205,"port")
	flag.Parse()

	addr := os.Args[1]

	config := &ssh.ClientConfig{
		Config: ssh.Config{},
		User:   user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return error(nil)
		},
		BannerCallback:    nil,
		ClientVersion:     "",
		HostKeyAlgorithms: nil,
		Timeout:           0,
	}

	conn, err := ssh.Dial("tcp", addr, config)
	handleError_(err, true)

	defer conn.Close()

	session, err := conn.NewSession()
	handleError_(err, false)

	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		handleError_(session.Close(), false)
		handleError_(err, true)
	}

	in, err := session.StdinPipe()
	handleError_(err, true)

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	handleError_(session.Shell(), true)

	scanner := bufio.NewScanner(os.Stdin)

	var current string

	for scanner.Scan() {
		current = scanner.Text()
		_, err = fmt.Fprintf(in, "%s\n", current)
		if current == "exit" {
			break
		}
	}

	handleError_(session.Wait(), false)
}

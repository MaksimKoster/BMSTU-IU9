package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os/exec"
	"strings"

	"github.com/gliderlabs/ssh"
	"setiLabs/config"
)

func main() {
	ssh.Handle(Handler)
	_ = ssh.ListenAndServe(":2205", nil,
		ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
			return pass == config.GetSshPassword()
		}),
	)
}

func Handler(session ssh.Session) {
	term := terminal.NewTerminal(session, "> ")
	//term := bufio.NewReader(session)
	log.Println("user connected")
	for {
		lineB, err := term.ReadLine()
		line := string(lineB)
		log.Println(line)
		if err != nil {
			break
		}
		//response := line + ":"
		log.Println(line)
		//if response != "" {
		//	_, _ = term.Write(append([]byte(response), '\n'))
		//}
		parts := strings.Fields(string(line))
		if len(parts) > 1 {
			line = parts[0]
			parts = parts[1:]
		} else {
			parts = nil
		}
		var resB []byte
		if len(parts) > 0 {
			log.Printf("exec %s + %s", line, parts)
			//log.Println(parts)
			resB, err = exec.Command(line, parts...).Output()
		} else {
			log.Printf("exec %s\n", line)
			resB, err = exec.Command(line).Output()
		}
		if err != nil {
			log.Println(err)
		}
		res := string(resB)
		if res != "" {
			log.Println("res:", res)
			b, err := fmt.Fprintf(session, "%s", res)
			if err != nil {
				log.Println(b, err)
			}
			//_, _ = term.Write(append(resB, '\n'))
		}
	}
	log.Println("terminal closed")
}

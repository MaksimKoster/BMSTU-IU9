package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"setiLabs/config"
)

func main() {
	//var ml config.Mail
	ml := config.GetMail("")
	from := mail.Address{Name: "Andrey", Address: "and.vladislavov@mail.ru"}
	to := mail.Address{Name: "Andrey", Address: "and.vladislavov@gmail.com"}
	subj := "Test SMTP msg"
	body := "Test SMTP msg"

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	serverName := ml.Addr
	host, port, _ := net.SplitHostPort(serverName)
	fmt.Println(fmt.Sprintf("%s:%s", host, port))

	auth := smtp.PlainAuth("", ml.Mail, ml.Pwd, host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	conn, err := tls.Dial("tcp", serverName, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}
	defer c.Quit()

	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}
	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}
}

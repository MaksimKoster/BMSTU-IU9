package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/jlaffaye/ftp"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
)

func main() {
	args := os.Args
	var port = "2105"
	var host = "185.20.227.83"
	for _, a := range args {
		if len(a) == 4 {
			port = a
		} else if a != "-p" {
			host = a
		}
	}
	client, err := ftp.Dial(fmt.Sprintf("%s:%s", host, port))

	if err != nil {
		log.Panic(err)
	}
	term := bufio.NewReader(os.Stdin)
	fmt.Printf("Connected to %s.\n220 Welcome to the Go FTP Server\n", host)
	fmt.Printf("Name (%s:%s): ", host, os.Getenv("USER"))

	li, _, _ := term.ReadLine()
	fmt.Printf("331 User name ok, password required\nPassword:")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	//pw, _, _ := term.ReadLine()
	err = client.Login(string(li), string(bytePassword))
	if err != nil {
		log.Panic(err)
	}
	defer client.Quit()
	fmt.Print("\n230 Password ok, continue\n")
	if err = client.ChangeDir("AVL"); err != nil {
		log.Panic(err)
	}
	cDir, _ := client.CurrentDir()
	fmt.Println(cDir)
	content, _ := client.NameList("")
	for _, s := range content {
		fmt.Println(s)
	}
L:
	for {
		fmt.Print("ftp> ")
		s, _, err := term.ReadLine()
		cmd := strings.Split(string(s), " ")
		if err != nil {
			break
		}
		switch cmd[0] {
		case "?", "help":
			getHelp()
		case "send", "put":
			if len(cmd) < 2 {
				fmt.Printf("(local-file) ")
				file, _, _ := term.ReadLine()
				cmd = append(cmd, string(file))
				fmt.Printf("(remote-file) ")
				file, _, _ = term.ReadLine()
				if str := string(file); str != "" {
					cmd = append(cmd, str)
				} else {
					fmt.Printf("usage: send local-file remote-file")
				}
			}
			if err = uploadFile(client, cmd[1:]); err != nil {
				log.Println(err)
			}
		case "get", "recv":
			if err = downloadFile(client); err != nil {
				log.Println(err)
			}
		case "mkdir":
			if len(cmd) < 2 {
				fmt.Printf("(directory-name) ")
				dir, _, _ := term.ReadLine()
				cmd = append(cmd, string(dir))
			}
			if err = makeDir(client, cmd[1]); err != nil {
				log.Println(err)
			}
		case "rm":
			if len(cmd) < 2 {
				fmt.Printf("(directory-name) ")
				dir, _, _ := term.ReadLine()
				cmd = append(cmd, string(dir))
			}
			if err = deleteFile(client, cmd[1]); err != nil {
				log.Println(err)
			}
		case "ls", "dir":
			if err = getContent(client); err != nil {
				log.Println(err)
			}
		case "cd":
			if len(cmd) < 2 {
				fmt.Printf("(remote-directory) ")
				dir, _, _ := term.ReadLine()
				cmd = append(cmd, string(dir))
			}
			if err = changeDir(client, cmd[1]); err != nil {
				log.Println(err)
			}
		case "size":
			if len(cmd) < 2 {
				fmt.Printf("(remote-file) ")
				dir, _, _ := term.ReadLine()
				cmd = append(cmd, string(dir))
			}
			if size, err := client.FileSize(cmd[1]); err != nil {
				log.Println(err)
			} else {
				log.Printf("213 %d", size)
			}
		case "rename":
			if err = client.Rename(cmd[1], cmd[2]); err != nil {
				log.Println(err)
			}
		case "!", "quit", "exit":
			fmt.Printf("221 Goodbye")
			break L
		default:
			fmt.Println("?Invalid command")
		}
	}
}

func changeDir(client *ftp.ServerConn, path string) error {
	if err := client.ChangeDir(path); err != nil {
		return err
	} else {
		fmt.Printf("250 Directory changed to /%s\n", path)
	}
	return nil
}

func uploadFile(client *ftp.ServerConn, filePath []string) error {
	data, err := ioutil.ReadFile(filePath[0])
	toStore := bytes.NewBufferString(string(data))
	if len(filePath) == 1 {
		err = client.Stor(filePath[0], toStore)
	} else {
		err = client.Stor(filePath[1], toStore)
	}
	if err != nil {
		return err
	}
	fmt.Println("SAVED")
	return nil
}

func downloadFile(client *ftp.ServerConn) error {
	var toRead string
	_, _ = fmt.Scanf("%s", &toRead)
	r, err := client.Retr(toRead)
	if err != nil {
		return err
	}
	file, err := os.Create(toRead)
	if err != nil {
		return err
	}
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if _, err = file.Write(buf); err != nil {
		return err
	}
	fmt.Println("DOWNLOADED")
	return nil
}

func deleteFile(client *ftp.ServerConn, toDelete string) error {
	err := client.Delete(toDelete)
	if err != nil {
		return err
	}
	fmt.Println("DELETED")
	return nil
}

func makeDir(client *ftp.ServerConn, name string) error {
	//var name string
	//_, _ = fmt.Scanf("%s", &name)
	err := client.MakeDir(name)
	if err != nil {
		return err
	}
	fmt.Println("DIR CREATED")
	return nil
}

func getContent(client *ftp.ServerConn) error {
	content, err := client.NameList("")
	if err != nil {
		return err
	}
	for _, s := range content {
		fmt.Printf("%s\n", s)
	}
	return nil
}

func getHelp() {
	file, err := ioutil.ReadFile("help.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(file))
}

/*
	- загрузку файла go ftp клиентом на ftp сервер;
	- скачивание файла go ftp клиентом с ftp сервера;
	- создание директории go ftp клиентом на ftp сервере;
	- удаление go ftp клиентом  файла на ftp сервере;
	- получение содержимого директории на ftp сервере с помощью go ftp клиента.
*/

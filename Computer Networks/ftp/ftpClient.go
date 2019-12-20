package main

import (
	"bytes"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	client, err := ftp.Dial("students.yss.su:21")

	if err != nil {
		log.Panic(err)
	}
	err = client.Login("ftpiu8", "3Ru7yOTA")
	if err != nil {
		log.Panic(err)
	}
	defer client.Quit()
	if err = client.ChangeDir("AVL"); err != nil {
		log.Panic(err)
	}
	cDir, _ := client.CurrentDir()
	fmt.Println(cDir)
	content, _ := client.NameList("")
	for _, s := range content {
		fmt.Println(s)
	}

	var inp string
	for {
		_, err = fmt.Scanf("%s", &inp)
		if err != nil {
			break
		}
		switch inp {
		case "help":
			GetHelp()
		case "upload":
			if err = UploadFile(client); err != nil {
				log.Println(err)
			}
		case "download":
			if err = DownloadFile(client); err != nil {
				log.Println(err)
			}
		case "mkdir":
			if err = MakeDir(client); err != nil {
				log.Println(err)
			}
		case "rm":
			if err = DeleteFile(client); err != nil {
				log.Println(err)
			}
		case "ls":
			if err = GetContent(client); err != nil {
				log.Println(err)
			}
		default:
			fmt.Println("Команда не распознана, попробуйте еще раз")
		}
	}
}

func UploadFile(client *ftp.ServerConn) error {
	var filePath string
	_, _ = fmt.Scanf("%s", &filePath)
	data, err := ioutil.ReadFile(filePath)
	toStore := bytes.NewBufferString(string(data))
	err = client.Stor(filePath, toStore)
	if err != nil {
		return err
	}
	fmt.Println("SAVED")
	return nil
}

func DownloadFile(client *ftp.ServerConn) error {
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

func DeleteFile(client *ftp.ServerConn) error {
	var toDelete string
	_, _ = fmt.Scanf("%s ", &toDelete)
	err := client.Delete(toDelete)
	if err != nil {
		return err
	}
	fmt.Println("DELETED")
	return nil
}

func MakeDir(client *ftp.ServerConn) error {
	var name string
	_, _ = fmt.Scanf("%s", &name)
	err := client.MakeDir(name)
	if err != nil {
		return err
	}
	fmt.Println("DIR CREATED")
	return nil
}

func GetContent(client *ftp.ServerConn) error {
	content, err := client.NameList("")
	if err != nil {
		return err
	}
	for _, s := range content {
		fmt.Printf("\t%s\n", s)
	}
	return nil
}

func GetHelp() {
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

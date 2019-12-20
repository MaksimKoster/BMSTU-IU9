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
	client, err := ftp.Dial("localhost:2121")

	if err != nil {
		log.Panic(err)
	}
	err = client.Login("admin", "123456")
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
			GetHelp_()
		case "upload":
			if err = UploadFile_(client); err != nil {
				log.Println(err)
			}
		case "download":
			if err = DownloadFile_(client); err != nil {
				log.Println(err)
			}
		case "mkdir":
			if err = MakeDir_(client); err != nil {
				log.Println(err)
			}
		case "rmdir":
			if err = RemoveDir(client); err != nil {
				log.Println(err)
			}
		case "rm":
			if err = DeleteFile_(client); err != nil {
				log.Println(err)
			}
		case "ls":
			if err = GetContent_(client); err != nil {
				log.Println(err)
			}
		default:
			fmt.Println("Команда не распознана, попробуйте еще раз")
		}
	}
}

func UploadFile_(client *ftp.ServerConn) error {
	var filePath, fileName string
	_, _ = fmt.Scanf("%s%s", &filePath, &fileName)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {

	}
	toStore := bytes.NewBufferString(string(data))
	err = client.Stor(fmt.Sprintf("%s/%s", filePath, fileName), toStore)
	if err != nil {
		return err
	}
	fmt.Println("SAVED")
	return nil
}

func DownloadFile_(client *ftp.ServerConn) error {
	var filePath, fileName string
	_, _ = fmt.Scanf("%s%s", &filePath, &fileName)
	r, err := client.Retr(fmt.Sprintf("%s/%s", filePath, fileName))
	if err != nil {
		return err
	}
	file, err := os.Create(fileName)
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

func DeleteFile_(client *ftp.ServerConn) error {
	var toDelete string
	_, _ = fmt.Scanf("%s", &toDelete)
	err := client.Delete(toDelete)
	if err != nil {
		return err
	}
	fmt.Println("DELETED")
	return nil
}

func MakeDir_(client *ftp.ServerConn) error {
	var name string
	_, _ = fmt.Scanf("%s", &name)
	err := client.MakeDir(name)
	if err != nil {
		return err
	}
	fmt.Println("DIR CREATED")
	return nil
}

func RemoveDir(client *ftp.ServerConn) error {
	var name string
	_, _ = fmt.Scanf("%s", &name)
	err := client.RemoveDir(name)
	if err != nil {
		return err
	}
	fmt.Println("DIR REMOVED")
	return nil
}

func GetContent_(client *ftp.ServerConn) error {
	content, err := client.NameList("")
	if err != nil {
		return err
	}
	for _, s := range content {
		fmt.Printf("\t%s\n", s)
	}
	return nil
}

func GetHelp_() {
	file, err := ioutil.ReadFile("help2.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(file))
}

/*
	- создание директории на удаленном go ftp сервере из go ftp клиента; (mkdir)
	- удаление директории на удаленном go ftp сервере из go ftp клиента; (rmdir)
	- передачу файла на удаленный go ftp сервер из текущей директории с go ftp 	клиента в заданную директорию go ftp сервера; (upload)
	- прием файла с удаленного go ftp сервера из заданной директории на go ftp 	клиент в текущую директорию.(download)
*/

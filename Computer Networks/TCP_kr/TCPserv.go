package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

// требуется только ниже для обработки примера

func main() {

	fmt.Println("Launching server...")

	// Устанавливаем прослушивание порта
	ln, _ := net.Listen("tcp", ":8084")

	// Открываем порт
	conn, _ := ln.Accept()

	// Запускаем цикл
	for {
		// Будем прослушивать все сообщения разделенные \n
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// Распечатываем полученое сообщение
		fmt.Print("Message Received:", message)
		arr := strings.Split(message, " ")
		max := strconv.Itoa(getMax(arr))
		log.Println(max)
		// Отправить новую строку обратно клиенту
		conn.Write([]byte(max + "\n"))
	}
}

func getMax(arr []string) int {
	max, _ := strconv.Atoi(arr[0])
	for _, v := range arr[1:] {
		cur, _ := strconv.Atoi(v)
		log.Println(cur, max)
		if cur > max {
			max = cur
		}
	}
	return max
}

// ВАРИАНТ 17 - макс число

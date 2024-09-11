package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

// Код клиента
func main() {
	// подключиться к серверу
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// запустить горутину, которая будет читать все сообщения от сервера и выводить их в консоль
	go clientReader(conn)

	// читать сообщения от stdin и отправлять их на сервер
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		data := scanner.Text()
		_, err := conn.Write([]byte(data))
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("reading standard input:", err)
	}
}

// clientReader выводит на экран все сообщения от сервера
func clientReader(conn net.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}

		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
}

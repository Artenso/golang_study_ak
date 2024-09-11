package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type client struct {
	conn net.Conn
	name string
	ch   chan<- string
}

var (
	// канал для всех входящих клиентов
	entering = make(chan client)
	// канал для сообщения о выходе клиента
	leaving = make(chan client)
	// канал для всех сообщений
	messages = make(chan string)
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

// broadcaster рассылает входящие сообщения всем клиентам
// следит за подключением и отключением клиентов
func broadcaster() {
	// здесь хранятся все подключенные клиенты
	clients := make(map[client]bool)

	for {

		msg := <-messages

		select {
		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			clients[cli] = false
		default:
		}

		for cli, connected := range clients {
			if connected {
				_, err := cli.conn.Write([]byte(msg))
				if err != nil {
					fmt.Printf("failed to send message to %s: %s\n", cli.name, err)
					continue
				}
			}
		}

	}

}

// handleConn обрабатывает входящие сообщения от клиента
func handleConn(conn net.Conn) {
	defer conn.Close()

	ch := make(chan string)

	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()

	cli := client{conn, who, ch}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			messages <- who + " has left"
			leaving <- cli
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}

		if n > 0 {
			messages <- who + ":" + string(buf[:n])
		}
	}
}

// clientWriter отправляет сообщения текущему клиенту
func clientWriter(conn net.Conn, ch <-chan string) {
	msg := <-ch

	_, err := conn.Write([]byte(msg))
	if err != nil {
		fmt.Printf("failed to send message: %s\n", err)
	}
}

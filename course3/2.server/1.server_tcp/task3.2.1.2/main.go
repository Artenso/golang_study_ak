package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}

	req := string(buf[:n])
	reqParts := strings.Split(req, " ")

	resp := fmt.Sprintf(
		"%s 404 Not Found\n%s\n\n%s",
		reqParts[2],
		"Content-Type: text/html",
		"<!DOCTYPE html>\n<html>\n<body>\n404 Not Found\n</body>\n</html>",
	)

	if strings.HasPrefix(req, "GET / ") {
		resp = fmt.Sprintf(
			"%s 200 OK\n%s\n\n%s",
			reqParts[2],
			"Content-Type: text/html",
			"<!DOCTYPE html>\n<html>\n<head>\n<title>Webserver</title>\n</head>\n<body>\nhello world\n</body>\n</html>",
		)
	}

	conn.Write([]byte(resp))

	defer conn.Close()
}

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}

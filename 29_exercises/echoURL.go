package main 

import (
	"net"
	"bufio"
	"io"
	"log"
	"fmt"
	"strings"
)

func handleConn (conn net.Conn) {
	defer conn.Close()
	
	i := 0
	var url string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		
		if i == 0 {
			url = strings.Fields(ln)[1]
			fmt.Println("URL ", url)
		} else {
			if ln == "" {
				break
			}
		}
		
		i++
	}
	
	//response
	body := url
	
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func main () {
	server, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer server.Close()
	
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handleConn(conn)
	}
}
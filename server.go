package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

var port = "0.0.0.0:9001"

func echoMessage(conn net.Conn) {

	r := bufio.NewReader(conn)

	for {

		message, err := r.ReadBytes(byte('\n'))

		switch err {

		case nil:

			break

		case io.EOF:

		default:

			fmt.Println("Error ", err)

		}
		conn.Write([]byte("From Server :"))
		conn.Write(message)
    //time and Ip Address
		fmt.Println("time stamp And IP Address : ", string(message), time.Now(), conn.RemoteAddr())

	}

}

var NumberOfConnections int = 0

func main() {

	fmt.Println("Start the server ")

	ln, err := net.Listen("tcp", port)

	for {

		conn, _ := ln.Accept()


//Number of connections
		NumberOfConnections++

		fmt.Println("Total Connected  Clients are:", NumberOfConnections)

	

		if err != nil {

			fmt.Println("Error ", err)

			continue

		}

		go echoMessage(conn)

	}

}

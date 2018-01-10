package main

import (
	"fmt"
	"net"
	"time"
	"bufio"
	"strconv"
	"monologued/rfc1288"
	"monologued/dotplan"
)
	

const PORT = 2003

func main() {
    server, err := net.Listen("tcp", ":" + strconv.Itoa(PORT))
    if server == nil {
		if (err != nil) {
			panic("couldn't start listening: " + err.Error())
		}
		panic("Couldn't start listening. Error undefined")
    }
	for {
		client, err := server.Accept()
		if client == nil {
			fmt.Printf("Connection request failed: %s\n", err.Error())
			continue
		}
		go Response(client)
	}
}


func Response(socket net.Conn) {
	defer func() {
		socket.Close()
	}()

	timeout := time.Second * 15
	socket.SetDeadline(time.Now().Add(timeout))
		
	buffer := bufio.NewReader(socket)
	line, _, err := buffer.ReadLine()

	if err != nil {
		socket.Write([]byte("400 Bad Request\r\n\r\n"))
		return
	}

	err, Request := rfc1288.ParseRfc1288Request(string(line))
	if err != nil {
		socket.Write([]byte("400 Bad Request\r\n\r\n"))
		return
	}

	if Request.Type == rfc1288.Remote {
		socket.Write([]byte("403 Forbidden - This server does not support remote requests\r\n\r\n"))
		return
	}

	if Request.Type == rfc1288.UserList {
		socket.Write([]byte("403 Forbidden - This server does not support user lists\r\n\r\n"))
		return
	}
	
	err, Data := dotplan.GetUserplan(Request.User)
	if err != nil {
		socket.Write([]byte(fmt.Sprintf("404 Not Found - No information for user '%s' found\r\n\r\n", *Request.User)))
		return
	}

	socket.Write(*Data)
}
	



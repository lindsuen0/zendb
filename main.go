// zendb - main.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package main

import (
	"bufio"
	"net"

	s "github.com/lindsuen0/zendb/stream"
	c "github.com/lindsuen0/zendb/util/config"
	d "github.com/lindsuen0/zendb/util/db"
	l "github.com/lindsuen0/zendb/util/log"
)

func init() {
	l.Setup()
	c.Setup()
	d.Setup()
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:"+c.DBConfig.Port)
	if err != nil {
		l.Logger.Fatalln("Error listening: ", err.Error())
	}
	defer listener.Close()

	l.Logger.Println("ZenDB server has been started. Listening on port " + c.DBConfig.Port + "...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			l.Logger.Fatalln("Error accepting connection: ", err.Error())
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// defer conn.Close()
	// // TODO
	// for {
	// 	reader := bufio.NewReader(conn)
	// 	var buf [128]byte
	// 	n, err := reader.Read(buf[:])
	// 	if err != nil {
	// 		fmt.Println("read from client failed, err: ", err)
	// 		break
	// 	}
	// 	recvStr := string(buf[:n])
	// 	fmt.Println("have recived msg from client: ", recvStr)
	// 	conn.Write([]byte(recvStr))
	// }
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [512]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			l.Logger.Println("An existing connection was closed by the remote host.")
			break
		}
		recvStr := string(buf[:n])
		l.Logger.Printf("Recived message: %q", recvStr)
		operatorTag := s.PreParseStruct(recvStr)
		if operatorTag == "0" {
			errOfParse := s.ParsePutStream(recvStr)
			if errOfParse != nil {
				l.Logger.Println(errOfParse)
			}
		} else if operatorTag == "1" {
			s.ParseDeleteStream(recvStr)
		} else if operatorTag == "2" {
			conn.Write([]byte(s.ParseGetStream(recvStr)))
		}
	}
}

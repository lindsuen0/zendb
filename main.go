// canodb - main.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"net"

	s "github.com/lindsuen/canodb/stream"
	c "github.com/lindsuen/canodb/util/config"
	d "github.com/lindsuen/canodb/util/db"
	l "github.com/lindsuen/canodb/util/log"
)

func init() {
	l.InitLog()
	c.InitLog()
	d.InitDB()
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:"+c.DBConfig.Port)
	if err != nil {
		l.Logger.Fatalln("Error listening: ", err.Error())
	}
	defer listener.Close()

	l.Logger.Println("CanoDB server has been started. Listening on port " + c.DBConfig.Port + "...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			l.Logger.Fatalln("Error accepting connection: ", err.Error())
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [512]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			l.Logger.Println("An existing connection was closed by the remote host.")
			break
		}
		recvByte := buf[:n]
		l.Logger.Printf("Recived message: %q", string(recvByte))
		operatorTag := s.PreParseMess(recvByte)
		if bytes.Equal(operatorTag, []byte("0")) {
			errOfParse := s.ParsePutMess(recvByte)
			if errOfParse != nil {
				l.Logger.Println(errOfParse)
			}
		} else if bytes.Equal(operatorTag, []byte("1")) {
			_ = s.ParseDelMess(recvByte)
		} else if bytes.Equal(operatorTag, []byte("2")) {
			b, _ := s.ParseGetMess(recvByte)
			conn.Write(b)
		}
	}
}

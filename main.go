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

	s "github.com/lindsuen0/canodb/stream"
	c "github.com/lindsuen0/canodb/util/config"
	d "github.com/lindsuen0/canodb/util/db"
	l "github.com/lindsuen0/canodb/util/log"
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
	// defer conn.Close()

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
		recvByte := buf[:n]
		l.Logger.Printf("Recived message: %q", string(recvByte))
		operatorTag := s.PreParseStruct(recvByte)
		if bytes.Equal(operatorTag, []byte("0")) {
			errOfParse := s.ParsePutStream(recvByte)
			if errOfParse != nil {
				l.Logger.Println(errOfParse)
			}
		} else if bytes.Equal(operatorTag, []byte("1")) {
			s.ParseDeleteStream(recvByte)
		} else if bytes.Equal(operatorTag, []byte("2")) {
			conn.Write(s.ParseGetStream(recvByte))
		}
	}
}

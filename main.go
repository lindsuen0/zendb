// zendb - main.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package main

import (
	"bufio"
	"log"
	"net"

	"github.com/lindsuen0/zendb/util/config"
	"github.com/lindsuen0/zendb/util/db"
)

func init() {
	config.Setup()
	db.Setup()
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:"+config.DBConfig.Port)
	if err != nil {
		log.Fatalln("Error listening: ", err.Error())
	}
	defer listener.Close()

	log.Println("ZenDB server has been started. Listening on port " + config.DBConfig.Port + "...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Error accepting connection: ", err.Error())
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [4096]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			log.Println("An existing connection was closed by the remote host.")
			break
		}
		recvStr := string(buf[:n])
		log.Println("Recived msg from client: , message: ", recvStr)
		conn.Write([]byte(recvStr))
	}
}

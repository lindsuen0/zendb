// zendb - main.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package main

import (
	"bufio"
	logg "github.com/lindsuen0/zendb/util/log"
	"log"
	"net"

	"github.com/lindsuen0/zendb/stream"
	"github.com/lindsuen0/zendb/util/config"
	"github.com/lindsuen0/zendb/util/db"
)

func init() {
	config.Setup()
	db.Setup()
	logg.Setup()
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:"+config.DBConfig.Port)
	if err != nil {
		log.Fatalln("[ZenDB] Error listening: ", err.Error())
	}
	defer listener.Close()

	log.Println("[ZenDB] ZenDB server has been started. Listening on port " + config.DBConfig.Port + "...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("[ZenDB] Error accepting connection: ", err.Error())
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
			logg.Println("An existing connection was closed by the remote host.")
			break
		}
		recvStr := string(buf[:n])
		log.Printf("[ZenDB] Recived message: %q", recvStr)
		errOfParse := stream.ParsePutStream(recvStr)
		if errOfParse != nil {
			log.Println(errOfParse)
		}
		// conn.Write([]byte(recvStr))
	}
}

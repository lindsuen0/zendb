// zendb - main.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package main

import (
	"bufio"
	"net"

	"github.com/lindsuen0/zendb/stream"
	"github.com/lindsuen0/zendb/util/config"
	"github.com/lindsuen0/zendb/util/db"
	logg "github.com/lindsuen0/zendb/util/log"
)

func init() {
	logg.Setup()
	config.Setup()
	db.Setup()

}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:"+config.DBConfig.Port)
	if err != nil {
		logg.Logger.Fatalln("Error listening: ", err.Error())
	}
	defer listener.Close()

	logg.Logger.Println("ZenDB server has been started. Listening on port " + config.DBConfig.Port + "...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			logg.Logger.Fatalln("Error accepting connection: ", err.Error())
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
			logg.Logger.Println("An existing connection was closed by the remote host.")
			break
		}
		recvStr := string(buf[:n])
		logg.Logger.Printf("Recived message: %q", recvStr)
		operatorTag := stream.PreParseStruct(recvStr)
		if operatorTag == "0" {
			errOfParse := stream.ParsePutStream(recvStr)
			if errOfParse != nil {
				logg.Logger.Println(errOfParse)
			}
		} else if operatorTag == "1" {
			logg.Logger.Println("1: Delete")
		} else if operatorTag == "2" {
			logg.Logger.Println("2: Get")
		}
	}
}

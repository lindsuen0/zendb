// zendb - client.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package client

import (
	"log"
	"net"

	"github.com/lindsuen0/zendb/stream"
)

type DB struct {
	ZenDBConnect net.Conn
}

func Connect(tcpAddress string) DB {
	conn, err := net.Dial("tcp", tcpAddress)
	if err != nil {
		log.Fatalln(err)
	}
	// defer conn.Close()

	// inputReader := bufio.NewReader(os.Stdin)
	// for {
	// 	input, _ := inputReader.ReadString('\n')
	// 	inputInfo := strings.Trim(input, "\r\n")
	// 	if strings.ToUpper(inputInfo) == "Q" {
	// 		return
	// 	}
	// 	_, err := conn.Write([]byte(inputInfo))
	// 	if err != nil {
	// 		return
	// 	}
	// 	buf := [512]byte{}
	// 	n, err := conn.Read(buf[:])
	// 	if err != nil {
	// 		log.Println("Recvied failed, err: ", err)
	// 		return
	// 	}
	// 	log.Println(string(buf[:n]))
	// }

	return DB{conn}
}

func (n *DB) Put(key string, value string) {
	s := stream.GeneratePutStream(key, value)
	operatorString := s.Operator.StartTag + s.Operator.OperatorContent + s.Operator.EndTag
	keyString := s.Key.StartTag + s.Key.KeyContent + s.Key.EndTag
	valueString := s.Value.StartTag + s.Value.ValueContent + s.Value.EndTag
	_, err := (*n).ZenDBConnect.Write([]byte(operatorString + keyString + valueString))
	if err != nil {
		log.Println(err)
	}
	// log.Println(operatorString + keyString + valueString)
	// stream.ParsePutStream(operatorString + keyString + valueString)
}

func (n *DB) Delete(key string) {
	stream.GenerateDeleteStream(key)
}

func (n *DB) Get(key string) string {
	stream.GenerateGetStream(key)
	return ""
}

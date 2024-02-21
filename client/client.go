// zendb - client.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package client

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"

	"github.com/lindsuen0/zendb/stream"
)

func Connect(tcpAddress string) {
	conn, err := net.Dial("tcp", tcpAddress)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Println("Recvied failed, err: ", err)
			return
		}
		log.Println(string(buf[:n]))
	}
}

func Put(key string, value string) error {
	s := stream.GeneratePutStream(key, value)
	operatorString := s.Operator.StartTag + s.Operator.OperatorContent + s.Operator.EndTag
	keyString := s.Key.StartTag + s.Key.KeyContent + s.Key.EndTag
	valueString := s.Value.StartTag + s.Value.ValueContent + s.Value.EndTag
	return stream.ParsePutStream(operatorString + keyString + valueString)
}

func Delete(key string) {
	stream.GenerateDeleteStream(key)
}

func Get(key string) string {
	stream.GenerateGetStream(key)
	return ""
}

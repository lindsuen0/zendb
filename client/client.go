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
)

type stream struct {
	operator operatorStruct
	key      keyStruct
	value    valueStruct
}

type operatorStruct struct {
	startTag        string
	operatorContent uint8
	endTag          string
}

type keyStruct struct {
	startTag   string
	keyContent string
	endTag     string
}

type valueStruct struct {
	startTag     string
	valueContent string
	endTag       string
}

func (b *operatorStruct) setOperatorTag() {
	b.startTag = ":"
	b.endTag = "\n"
}

func (b *keyStruct) setKeyTag() {
	b.startTag = "$"
	b.endTag = "\n"
}

func (b *valueStruct) setValueTag() {
	b.startTag = "-"
	b.endTag = "\n"
}

func (b *operatorStruct) setOperatorContent(s uint8) {
	b.operatorContent = s
}

func (b *keyStruct) setKeyContent(s string) {
	b.keyContent = s
}

func (b *valueStruct) setValueContent(s string) {
	b.valueContent = s
}

// setPutStream
// 0: Put, 1: Delete
func setPutStream(key string, value string) stream {
	stream := new(stream)
	stream.operator.setOperatorTag()
	stream.key.setKeyTag()
	stream.value.setValueTag()
	stream.operator.setOperatorContent(0)
	stream.key.setKeyContent(key)
	stream.value.setValueContent(value)
	return *stream
}

// GenerateDeleteStream
// 0: Put, 1: Delete
//
//	func GenerateDeleteStream(key string) stream {
//		stream := new(stream)
//		stream.operator.setOperatorTag()
//		stream.key.setKeyTag()
//		stream.value.setValueTag()
//		stream.operator.setOperatorContent(1)
//		stream.key.setKeyContent(key)
//		stream.value.setValueContent("")
//
//		return *stream
//	}

func Connect(add string) {
	conn, err := net.Dial("tcp", add)
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

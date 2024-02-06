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

type Stream struct {
	operator OperatorStruct
	key      KeyStruct
	value    ValueStruct
}

type OperatorStruct struct {
	startTag        string
	operatorContent string
	endTag          string
}

type KeyStruct struct {
	startTag   string
	keyContent string
	endTag     string
}

type ValueStruct struct {
	startTag     string
	valueContent string
	endTag       string
}

func (b *OperatorStruct) setOperatorTag() {
	b.startTag = ":"
	b.endTag = "\\n"
}

func (b *KeyStruct) setKeyTag() {
	b.startTag = "$"
	b.endTag = "\\n"
}

func (b *ValueStruct) setValueTag() {
	b.startTag = "-"
	b.endTag = "\\n"
}

func (b *OperatorStruct) setOperatorContent(s string) {
	b.operatorContent = s
}

func (b *KeyStruct) setKeyContent(s string) {
	b.keyContent = s
}

func (b *ValueStruct) setValueContent(s string) {
	b.valueContent = s
}

// SetPutStream
// 0: Put, 1: Delete
// stream:
// stream.operator.startTag+stream.operator.operatorContent+stream.operator.endTag
// stream.key.startTag+stream.key.keyContent+stream.key.endTag
// stream.value.startTag+stream.value.valueContent+stream.value.endTag
func SetPutStream(key string, value string) Stream {
	stream := new(Stream)
	stream.operator.setOperatorTag()
	stream.key.setKeyTag()
	stream.value.setValueTag()
	stream.operator.setOperatorContent("0")
	stream.key.setKeyContent(key)
	stream.value.setValueContent(value)
	return *stream
}

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

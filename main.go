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

	m "github.com/lindsuen/canodb/message"
	c "github.com/lindsuen/canodb/util/config"
	d "github.com/lindsuen/canodb/util/db"
	l "github.com/lindsuen/canodb/util/log"
)

// :VALUE\n
// ^STATUSCODE\n
//
// STATUSCODE:
//
//	20 SUCCESS, 30 NOTFOUND, 40 EXCEPTION
type sMessage struct {
	startTag []byte
	Content  []byte
	endTag   []byte
}

func (m *sMessage) setValueTag() {
	m.startTag = []byte(":")
	m.endTag = []byte("\n")
}

func (m *sMessage) setStatusCodeTag() {
	m.startTag = []byte("^")
	m.endTag = []byte("\n")
}

func (m *sMessage) setContent(c []byte) {
	m.Content = c
}

func init() {
	l.InitLog()
	c.InitConfig()
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
		resMessage := new(sMessage)
		reader := bufio.NewReader(conn)
		var buf [512]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			l.Logger.Println("An existing connection was closed by the remote host.")
			break
		}
		recvByte := buf[:n]
		l.Logger.Printf("Recived message: %q", string(recvByte))
		operatorTag := m.PreParseMess(recvByte)
		if bytes.Equal(operatorTag, []byte("0")) {
			errOfParse := m.ParsePutMess(recvByte)
			if errOfParse != nil {
				l.Logger.Println(errOfParse)
			} else {
				resMessage.setStatusCodeTag()
				resMessage.setContent([]byte("20"))
				conn.Write(mergeByteSlice(resMessage.startTag, resMessage.Content, resMessage.endTag))
			}
		} else if bytes.Equal(operatorTag, []byte("1")) {
			_ = m.ParseDelMess(recvByte)
			resMessage.setStatusCodeTag()
			resMessage.setContent([]byte("20"))
			conn.Write(mergeByteSlice(resMessage.startTag, resMessage.Content, resMessage.endTag))
		} else if bytes.Equal(operatorTag, []byte("2")) {
			b, _ := m.ParseGetMess(recvByte)
			if byteSliceIsNil(b) {
				resMessage.setStatusCodeTag()
				resMessage.setContent([]byte("30"))
				conn.Write(mergeByteSlice(resMessage.startTag, resMessage.Content, resMessage.endTag))
			} else {
				resMessage.setValueTag()
				resMessage.setContent(b)
				conn.Write(mergeByteSlice(resMessage.startTag, resMessage.Content, resMessage.endTag))
			}
		}
	}
}

func mergeByteSlice(startTag []byte, content []byte, endTag []byte) []byte {
	return append(append(startTag, content...), endTag...)
}

func byteSliceIsNil(b []byte) bool {
	var tempSlice []byte
	return bytes.Equal(b, tempSlice)
}

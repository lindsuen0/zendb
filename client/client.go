// canodb - client.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package client

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"time"
)

type Driver struct {
	Conn net.Conn
}

type message struct {
	operator operatorStruct
	key      keyStruct
	value    valueStruct
}

type operatorStruct struct {
	startTag        []byte
	operatorContent []byte
	endTag          []byte
}

type keyStruct struct {
	startTag   []byte
	keyContent []byte
	endTag     []byte
}

type valueStruct struct {
	startTag     []byte
	valueContent []byte
	endTag       []byte
}

func (b *operatorStruct) setOperatorTag() {
	b.startTag = []byte(":")
	b.endTag = []byte("\n")
}

func (b *keyStruct) setKeyTag() {
	b.startTag = []byte("$")
	b.endTag = []byte("\n")
}

func (b *valueStruct) setValueTag() {
	b.startTag = []byte("-")
	b.endTag = []byte("\n")
}

func (b *operatorStruct) setOperatorContent(s []byte) {
	b.operatorContent = s
}

func (b *keyStruct) setKeyContent(s []byte) {
	b.keyContent = s
}

func (b *valueStruct) setValueContent(s []byte) {
	b.valueContent = s
}

const (
	READTIMEOUT = 300 * time.Millisecond
)

func createPutMess(key []byte, value []byte) message {
	putMessage := new(message)
	putMessage.operator.setOperatorTag()
	putMessage.key.setKeyTag()
	putMessage.value.setValueTag()
	putMessage.operator.setOperatorContent([]byte("0"))
	putMessage.key.setKeyContent(key)
	putMessage.value.setValueContent(value)
	return *putMessage
}

func createDelMess(key []byte) message {
	delMessage := new(message)
	delMessage.operator.setOperatorTag()
	delMessage.key.setKeyTag()
	delMessage.operator.setOperatorContent([]byte("1"))
	delMessage.key.setKeyContent(key)
	return *delMessage
}

func createGetMess(key []byte) message {
	getMessage := new(message)
	getMessage.operator.setOperatorTag()
	getMessage.key.setKeyTag()
	getMessage.operator.setOperatorContent([]byte("2"))
	getMessage.key.setKeyContent(key)
	return *getMessage
}

func Connect(tcpAddress string) (*Driver, error) {
	conn, err := net.Dial("tcp", tcpAddress)
	return &Driver{conn}, err
}

func (n *Driver) Close() error {
	return n.Conn.Close()
}

func (n *Driver) Put(key []byte, value []byte) error {
	p := createPutMess(key, value)
	operatorString := mergeByteSlice(p.operator.startTag, p.operator.operatorContent, p.operator.endTag)
	keyString := mergeByteSlice(p.key.startTag, p.key.keyContent, p.key.endTag)
	valueString := mergeByteSlice(p.value.startTag, p.value.valueContent, p.value.endTag)

	_, err := n.Conn.Write(append(append(operatorString, keyString...), valueString...))
	return err
}

func (n *Driver) Delete(key []byte) error {
	p := createDelMess(key)
	operatorString := mergeByteSlice(p.operator.startTag, p.operator.operatorContent, p.operator.endTag)
	keyString := mergeByteSlice(p.key.startTag, p.key.keyContent, p.key.endTag)

	_, err := n.Conn.Write(append(operatorString, keyString...))
	return err
}

func (n *Driver) Get(key []byte) ([]byte, error) {
	var err error
	p := createGetMess(key)
	operatorString := mergeByteSlice(p.operator.startTag, p.operator.operatorContent, p.operator.endTag)
	keyString := mergeByteSlice(p.key.startTag, p.key.keyContent, p.key.endTag)

	_, err = n.Conn.Write(append(operatorString, keyString...))
	if err != nil {
		fmt.Println(err)
	}

	_ = n.Conn.SetReadDeadline(time.Now().Add(READTIMEOUT))
	var buf [256]byte
	b, _ := n.Conn.Read(buf[:])
	if bytes.Equal(buf[:1], []byte(":")) {
		return parseMess(buf[:b], ":", "\n"), nil
	} else if bytes.Equal(buf[:1], []byte("^")) {
		statusCode := parseMess(buf[:b], "^", "\n")
		if bytes.Equal(statusCode, []byte("30")) {
			return nil, errors.New("canodb: not found")
		}
	}
	return nil, errors.New("something error")
}

func mergeByteSlice(startTag []byte, content []byte, endTag []byte) []byte {
	return append(append(startTag, content...), endTag...)
}

// func byteSliceIsNil(b []byte) bool {
// 	var tempSlice []byte
// 	return bytes.Equal(b, tempSlice)
// }

func parseMess(message []byte, startTag string, endTag string) []byte {
	var startTagIndex, endTagIndex int

	for k, v := range message {
		if string(v) == startTag {
			startTagIndex = k
			break
		}
	}

	tempIndex := startTagIndex
	for ; ; tempIndex++ {
		if string(message[tempIndex]) == endTag {
			endTagIndex = tempIndex
			break
		}
	}

	return message[startTagIndex+1 : endTagIndex]
}

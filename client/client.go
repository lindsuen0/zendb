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

type Message struct {
	Operator operatorStruct
	Key      keyStruct
	Value    valueStruct
}

type operatorStruct struct {
	StartTag        []byte
	OperatorContent []byte
	EndTag          []byte
}

type keyStruct struct {
	StartTag   []byte
	KeyContent []byte
	EndTag     []byte
}

type valueStruct struct {
	StartTag     []byte
	ValueContent []byte
	EndTag       []byte
}

func (b *operatorStruct) setOperatorTag() {
	b.StartTag = []byte(":")
	b.EndTag = []byte("\n")
}

func (b *keyStruct) setKeyTag() {
	b.StartTag = []byte("$")
	b.EndTag = []byte("\n")
}

func (b *valueStruct) setValueTag() {
	b.StartTag = []byte("-")
	b.EndTag = []byte("\n")
}

func (b *operatorStruct) setOperatorContent(s []byte) {
	b.OperatorContent = s
}

func (b *keyStruct) setKeyContent(s []byte) {
	b.KeyContent = s
}

func (b *valueStruct) setValueContent(s []byte) {
	b.ValueContent = s
}

const (
	READTIMEOUT = 300 * time.Millisecond
)

func CreatePutMess(key []byte, value []byte) Message {
	putMessage := new(Message)
	putMessage.Operator.setOperatorTag()
	putMessage.Key.setKeyTag()
	putMessage.Value.setValueTag()
	putMessage.Operator.setOperatorContent([]byte("0"))
	putMessage.Key.setKeyContent(key)
	putMessage.Value.setValueContent(value)
	return *putMessage
}

func CreateDelMess(key []byte) Message {
	delMessage := new(Message)
	delMessage.Operator.setOperatorTag()
	delMessage.Key.setKeyTag()
	delMessage.Operator.setOperatorContent([]byte("1"))
	delMessage.Key.setKeyContent(key)
	return *delMessage
}

func CreateGetMess(key []byte) Message {
	getMessage := new(Message)
	getMessage.Operator.setOperatorTag()
	getMessage.Key.setKeyTag()
	getMessage.Operator.setOperatorContent([]byte("2"))
	getMessage.Key.setKeyContent(key)
	return *getMessage
}

func Connect(tcpAddress string) (*Driver, error) {
	conn, err := net.Dial("tcp", tcpAddress)
	return &Driver{conn}, err
}

func (n *Driver) Put(key []byte, value []byte) error {
	p := CreatePutMess(key, value)
	operatorString := mergeByteSlice(p.Operator.StartTag, p.Operator.OperatorContent, p.Operator.EndTag)
	keyString := mergeByteSlice(p.Key.StartTag, p.Key.KeyContent, p.Key.EndTag)
	valueString := mergeByteSlice(p.Value.StartTag, p.Value.ValueContent, p.Value.EndTag)

	_, err := n.Conn.Write(append(append(operatorString, keyString...), valueString...))
	return err
}

func (n *Driver) Delete(key []byte) error {
	p := CreateDelMess(key)
	operatorString := mergeByteSlice(p.Operator.StartTag, p.Operator.OperatorContent, p.Operator.EndTag)
	keyString := mergeByteSlice(p.Key.StartTag, p.Key.KeyContent, p.Key.EndTag)

	_, err := n.Conn.Write(append(operatorString, keyString...))
	return err
}

func (n *Driver) Get(key []byte) ([]byte, error) {
	var err error
	p := CreateGetMess(key)
	operatorString := mergeByteSlice(p.Operator.StartTag, p.Operator.OperatorContent, p.Operator.EndTag)
	keyString := mergeByteSlice(p.Key.StartTag, p.Key.KeyContent, p.Key.EndTag)

	_, err = n.Conn.Write(append(operatorString, keyString...))
	if err != nil {
		fmt.Println(err)
	}

	_ = n.Conn.SetReadDeadline(time.Now().Add(READTIMEOUT))
	buf := [256]byte{}
	b, _ := n.Conn.Read(buf[:])
	if byteSliceIsNil(buf[:b]) {
		err = errors.New("canodb: not found")
	}
	return buf[:b], err
}

func mergeByteSlice(startTag []byte, content []byte, endTag []byte) []byte {
	return append(append(startTag, content...), endTag...)
}

func byteSliceIsNil(b []byte) bool {
	var tempSlice []byte
	return bytes.Equal(b, tempSlice)
}

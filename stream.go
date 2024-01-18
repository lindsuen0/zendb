// zendb - stream.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

type Stream struct {
	operator operatorStruct
	key      keyStruct
	value    valueStruct
	client   string
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

// func (b *operatorStruct) setKey() {
//
// }

// GeneratePutStream
// 0: Put, 1: Delete
func GeneratePutStream(key string, value string) Stream {
	stream := new(Stream)
	stream.client, _ = os.Hostname()
	stream.operator.setOperatorTag()
	stream.key.setKeyTag()
	stream.value.setValueTag()
	stream.operator.setOperatorContent(0)
	stream.key.setKeyContent(key)
	stream.value.setValueContent(value)
	fmt.Println(stream)
	return *stream
}

// GenerateDeleteStream
// 0: Put, 1: Delete
func GenerateDeleteStream(key string) Stream {
	stream := new(Stream)
	stream.client, _ = os.Hostname()
	stream.operator.setOperatorTag()
	stream.key.setKeyTag()
	stream.value.setValueTag()
	stream.operator.setOperatorContent(1)
	stream.key.setKeyContent(key)
	stream.value.setValueContent("")

	return *stream
}

func ParsePutStream(s *Stream) {

}

// canodb - stream.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package stream

import (
	"log"

	d "github.com/lindsuen/canodb/util/db"
)

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

func CreatePutMess(key []byte, value []byte) Message {
	stream := new(Message)
	stream.Operator.setOperatorTag()
	stream.Key.setKeyTag()
	stream.Value.setValueTag()
	stream.Operator.setOperatorContent([]byte("0"))
	stream.Key.setKeyContent(key)
	stream.Value.setValueContent(value)
	return *stream
}

func CreateDelMess(key []byte) Message {
	stream := new(Message)
	stream.Operator.setOperatorTag()
	stream.Key.setKeyTag()
	stream.Operator.setOperatorContent([]byte("1"))
	stream.Key.setKeyContent(key)
	return *stream
}

func CreateGetMess(key []byte) Message {
	stream := new(Message)
	stream.Operator.setOperatorTag()
	stream.Key.setKeyTag()
	stream.Operator.setOperatorContent([]byte("2"))
	stream.Key.setKeyContent(key)
	return *stream
}

func PreParseMess(message []byte) []byte {
	return parseMess(message, ":", "\n")
}

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

// ParsePutMess parses the stream of PUT.
// operatorTag: 0: Put, 1: Delete, 2: Get
func ParsePutMess(m []byte) error {
	operatorTag := parseMess(m, ":", "\n")
	keyContent := parseMess(m, "$", "\n")
	valueContent := parseMess(m, "-", "\n")

	if string(operatorTag) != "0" {
		log.Println("parse put error")
	}

	return d.PutData(keyContent, valueContent)
}

// ParseDelMess parses the stream of DELETE.
// operatorTag: 0: Put, 1: Delete, 2: Get
func ParseDelMess(m []byte) error {
	operatorTag := parseMess(m, ":", "\n")
	keyContent := parseMess(m, "$", "\n")

	if string(operatorTag) != "1" {
		log.Println("parse delete error")
	}

	err := d.DeleteData(keyContent)
	return err
}

// ParseGetMess parses the stream of GET.
// operatorTag: 0: Put, 1: Delete, 2: Get
func ParseGetMess(m []byte) ([]byte, error) {
	operatorTag := parseMess(m, ":", "\n")
	keyContent := parseMess(m, "$", "\n")

	if string(operatorTag) != "2" {
		log.Println("parse get error")
	}

	s, err := d.GetData(keyContent)
	return s, err
}

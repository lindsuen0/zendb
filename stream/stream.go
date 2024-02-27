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

type Stream struct {
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

func GeneratePutStream(key []byte, value []byte) Stream {
	stream := new(Stream)
	stream.Operator.setOperatorTag()
	stream.Key.setKeyTag()
	stream.Value.setValueTag()
	stream.Operator.setOperatorContent([]byte("0"))
	stream.Key.setKeyContent(key)
	stream.Value.setValueContent(value)
	return *stream
}

func GenerateDeleteStream(key []byte) Stream {
	stream := new(Stream)
	stream.Operator.setOperatorTag()
	stream.Key.setKeyTag()
	stream.Operator.setOperatorContent([]byte("1"))
	stream.Key.setKeyContent(key)
	return *stream
}

func GenerateGetStream(key []byte) Stream {
	stream := new(Stream)
	stream.Operator.setOperatorTag()
	stream.Key.setKeyTag()
	stream.Operator.setOperatorContent([]byte("2"))
	stream.Key.setKeyContent(key)
	return *stream
}

func PreParseStruct(message []byte) []byte {
	return parseStruct(message, ":", "\n")
}

func parseStruct(message []byte, startTag string, endTag string) []byte {
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

// ParsePutStream parses the stream of PUT.
// operatorTag:
// 0: Put, 1: Delete, 2: Get
func ParsePutStream(m []byte) error {
	operatorTag := parseStruct(m, ":", "\n")
	keyContent := parseStruct(m, "$", "\n")
	valueContent := parseStruct(m, "-", "\n")

	if string(operatorTag) != "0" {
		log.Println("error")
	}

	return d.PutData(keyContent, valueContent)
}

// ParseDeleteStream parses the stream of DELETE.
// operatorTag:
// 0: Put, 1: Delete, 2: Get
func ParseDeleteStream(m []byte) {
	operatorTag := parseStruct(m, ":", "\n")
	keyContent := parseStruct(m, "$", "\n")

	if string(operatorTag) != "1" {
		log.Println("error")
	}

	d.DeleteData(keyContent)
}

// ParseGetStream parses the stream of GET.
// operatorTag:
// 0: Put, 1: Delete, 2: Get
func ParseGetStream(m []byte) []byte {
	operatorTag := parseStruct(m, ":", "\n")
	keyContent := parseStruct(m, "$", "\n")

	if string(operatorTag) != "2" {
		log.Println("error")
	}

	s, _ := d.GetData(keyContent)
	return s
}

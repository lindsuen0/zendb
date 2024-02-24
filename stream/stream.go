// zendb - stream.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package stream

import (
	"log"

	"github.com/lindsuen0/zendb/util/db"
)

type Stream struct {
	Operator operatorStruct
	Key      keyStruct
	Value    valueStruct
}

type operatorStruct struct {
	StartTag        string
	OperatorContent string
	EndTag          string
}

type keyStruct struct {
	StartTag   string
	KeyContent string
	EndTag     string
}

type valueStruct struct {
	StartTag     string
	ValueContent string
	EndTag       string
}

func (b *operatorStruct) setOperatorTag() {
	b.StartTag = ":"
	b.EndTag = "\n"
}

func (b *keyStruct) setKeyTag() {
	b.StartTag = "$"
	b.EndTag = "\n"
}

func (b *valueStruct) setValueTag() {
	b.StartTag = "-"
	b.EndTag = "\n"
}

func (b *operatorStruct) setOperatorContent(s string) {
	b.OperatorContent = s
}

func (b *keyStruct) setKeyContent(s string) {
	b.KeyContent = s
}

func (b *valueStruct) setValueContent(s string) {
	b.ValueContent = s
}

// GeneratePutStream
// operatorTag:
// 0: Put, 1: Delete, 2: Get
//
// stream:
// stream.Operator.StartTag+stream.Operator.OperatorContent+stream.Operator.EndTag
// stream.Key.StartTag+stream.Key.KeyContent+stream.Key.EndTag
// stream.Value.StartTag+stream.Value.ValueContent+stream.Value.EndTag
func GeneratePutStream(key string, value string) Stream {
	stream := new(Stream)
	stream.Operator.setOperatorTag()
	stream.Key.setKeyTag()
	stream.Value.setValueTag()
	stream.Operator.setOperatorContent("0")
	stream.Key.setKeyContent(key)
	stream.Value.setValueContent(value)
	return *stream
}

// GenerateDeleteStream
// operatorTag:
// 0: Put, 1: Delete, 2: Get
//
// stream:
// stream.Operator.StartTag+stream.Operator.OperatorContent+stream.Operator.EndTag
// stream.Key.StartTag+stream.Key.KeyContent+stream.Key.EndTag
func GenerateDeleteStream(key string) {
	stream := new(Stream)
	stream.Operator.setOperatorTag()
	stream.Key.setKeyTag()
	stream.Operator.setOperatorContent("1")
	stream.Key.setKeyContent(key)
}

// GenerateGetStream
// operatorTag:
// 0: Put, 1: Delete, 2: Get
//
// stream:
// stream.Operator.StartTag+stream.Operator.OperatorContent+stream.Operator.EndTag
// stream.Key.StartTag+stream.Key.KeyContent+stream.Key.EndTag
func GenerateGetStream(key string) {
	stream := new(Stream)
	stream.Operator.setOperatorTag()
	stream.Key.setKeyTag()
	stream.Operator.setOperatorContent("2")
	stream.Key.setKeyContent(key)
}

func parseStruct(message string, startTag string, endTag string) string {
	messageSlice := []byte(message)
	var startTagIndex, endTagIndex int

	for k, v := range messageSlice {
		if string(v) == startTag {
			startTagIndex = k
			break
		}
	}

	tempIndex := startTagIndex
	for ; ; tempIndex++ {
		if string(messageSlice[tempIndex]) == endTag {
			endTagIndex = tempIndex
			break
		}
	}

	return string(messageSlice[startTagIndex+1 : endTagIndex])
}

// ParsePutStream parses the stream of PUT.
// operatorTag:
// 0: Put, 1: Delete, 2: Get
func ParsePutStream(m string) error {
	operatorTag := parseStruct(m, ":", "\n")
	keyContent := parseStruct(m, "$", "\n")
	valueContent := parseStruct(m, "-", "\n")

	if operatorTag != "0" {
		log.Println("error")
	}

	return db.PutData(keyContent, valueContent)
}

// ParseDeleteStream parses the stream of DELETE.
// operatorTag:
// 0: Put, 1: Delete, 2: Get
func ParseDeleteStream(m string) {
	operatorTag := parseStruct(m, ":", "\n")
	keyContent := parseStruct(m, "$", "\n")
	valueContent := parseStruct(m, "-", "\n")

	if operatorTag != "1" {
		log.Println("error")
	}
	if valueContent != "" {
		log.Println("error")
	}

	db.DeleteData(keyContent)
}

// ParseGetStream parses the stream of GET.
// operatorTag:
// 0: Put, 1: Delete, 2: Get
func ParseGetStream(m string) string {
	operatorTag := parseStruct(m, ":", "\n")
	keyContent := parseStruct(m, "$", "\n")
	valueContent := parseStruct(m, "-", "\n")

	if operatorTag != "2" {
		log.Println("error")
	}
	if valueContent != "" {
		log.Println("error")
	}

	s, _ := db.GetData(keyContent)

	return s
}

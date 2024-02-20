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
	operator operatorStruct
	key      keyStruct
	value    valueStruct
}

type operatorStruct struct {
	startTag        string
	operatorContent string
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
	b.endTag = "\\n"
}

func (b *keyStruct) setKeyTag() {
	b.startTag = "$"
	b.endTag = "\\n"
}

func (b *valueStruct) setValueTag() {
	b.startTag = "-"
	b.endTag = "\\n"
}

func (b *operatorStruct) setOperatorContent(s string) {
	b.operatorContent = s
}

func (b *keyStruct) setKeyContent(s string) {
	b.keyContent = s
}

func (b *valueStruct) setValueContent(s string) {
	b.valueContent = s
}

// GeneratePutStream
// 0: Put, 1: Delete
// stream:
// stream.operator.startTag+stream.operator.operatorContent+stream.operator.endTag
// stream.key.startTag+stream.key.keyContent+stream.key.endTag
// stream.value.startTag+stream.value.valueContent+stream.value.endTag
func GeneratePutStream(key string, value string) Stream {
	stream := new(Stream)
	stream.operator.setOperatorTag()
	stream.key.setKeyTag()
	stream.value.setValueTag()
	stream.operator.setOperatorContent("0")
	stream.key.setKeyContent(key)
	stream.value.setValueContent(value)
	return *stream
}

// GenerateDeleteStream
// 0: Put, 1: Delete
// stream:
// stream.operator.startTag+stream.operator.operatorContent+stream.operator.endTag
// stream.key.startTag+stream.key.keyContent+stream.key.endTag
func GenerateDeleteStream(key string) {
	stream := new(Stream)
	stream.operator.setOperatorTag()
	stream.key.setKeyTag()
	stream.operator.setOperatorContent("1")
	stream.key.setKeyContent(key)
}

func parseStruct(message string, startTag string, endTag string) string {
	messageSlice := []rune(message)
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

// ParsePutStream
// 0: Put, 1: Delete
func ParsePutStream(m string) {
	operatorTag := parseStruct(m, ":", "\n")
	keyContent := parseStruct(m, "$", "\n")
	valueContent := parseStruct(m, "-", "\n")

	if operatorTag != "0" {
		log.Println("error")
	}

	db.PutData(keyContent, valueContent)
}

// ParseDeleteStream
// 0: Put, 1: Delete
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

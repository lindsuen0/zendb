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

// ParsePutMess parses the message of PUT.
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

// ParseDelMess parses the message of DELETE.
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

// ParseGetMess parses the message of GET.
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

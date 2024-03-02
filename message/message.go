// canodb - message.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package message

import (
	d "github.com/lindsuen/canodb/util/db"
	l "github.com/lindsuen/canodb/util/log"
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
// operatorContent: 0: Put, 1: Delete, 2: Get
func ParsePutMess(m []byte) error {
	operatorContent := parseMess(m, ":", "\n")
	keyContent := parseMess(m, "$", "\n")
	valueContent := parseMess(m, "-", "\n")

	if string(operatorContent) != "0" {
		l.Logger.Println("Unable to parse Put() message.")
	}

	return d.PutData(keyContent, valueContent)
}

// ParseDelMess parses the message of DELETE.
// operatorContent: 0: Put, 1: Delete, 2: Get
func ParseDelMess(m []byte) error {
	operatorContent := parseMess(m, ":", "\n")
	keyContent := parseMess(m, "$", "\n")

	if string(operatorContent) != "1" {
		l.Logger.Println("Unable to parse Delete() message.")
	}

	err := d.DeleteData(keyContent)
	return err
}

// ParseGetMess parses the message of GET.
// operatorContent: 0: Put, 1: Delete, 2: Get
func ParseGetMess(m []byte) ([]byte, error) {
	operatorContent := parseMess(m, ":", "\n")
	keyContent := parseMess(m, "$", "\n")

	if string(operatorContent) != "2" {
		l.Logger.Println("Unable to parse Get() message.")
	}

	s, err := d.GetData(keyContent)
	return s, err
}

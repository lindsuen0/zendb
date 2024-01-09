// zendb - stream.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package main

type byteStream struct {
	operator operatorStruct
	key      keyStruct
	value    valueStruct
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

func (b *operatorStruct) setTag() {
	b.startTag = ":"
	b.endTag = "\n"
}

func (b *keyStruct) setTag() {
	b.startTag = "$"
	b.endTag = "\n"
}

func (b *valueStruct) setTag() {
	b.startTag = "-"
	b.endTag = "\n"
}

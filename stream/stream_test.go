// zendb - stream_test.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package stream

import (
	"fmt"
	"testing"
)

func TestGeneratePutStream(t *testing.T) {
	stream1 := GeneratePutStream("key1", "value1")
	operatorStream1 := stream1.operator.startTag + stream1.operator.operatorContent + stream1.operator.endTag
	keyStream1 := stream1.key.startTag + stream1.key.keyContent + stream1.key.endTag
	valueStream1 := stream1.value.startTag + stream1.value.valueContent + stream1.value.endTag
	t.Log(operatorStream1 + keyStream1 + valueStream1)
}

func TestGenerateDeleteStream(t *testing.T) {
	GenerateDeleteStream("key1")
}

func TestParseStruct(t *testing.T) {
	s := ":0\n$KEY321\n-VALUE123\n"
	d := ParseStruct(s, "-", "\n")

	fmt.Println(d)
}

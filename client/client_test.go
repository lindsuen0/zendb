package client

import (
	"testing"
)

func TestSetPutStream(t *testing.T) {
	stream1 := SetPutStream("key1", "value1")
	operatorStream1 := stream1.operator.startTag + stream1.operator.operatorContent + stream1.operator.endTag
	keyStream1 := stream1.key.startTag + stream1.key.keyContent + stream1.key.endTag
	valueStream1 := stream1.value.startTag + stream1.value.valueContent + stream1.value.endTag
	t.Log(operatorStream1 + keyStream1 + valueStream1)
}

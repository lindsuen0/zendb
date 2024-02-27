package iterator_test

import (
	"testing"

	"github.com/lindsuen/canodb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}

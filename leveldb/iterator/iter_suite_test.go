package iterator_test

import (
	"testing"

	"github.com/lindsuen0/zendb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}

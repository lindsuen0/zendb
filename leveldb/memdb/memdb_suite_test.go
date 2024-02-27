package memdb

import (
	"testing"

	"github.com/lindsuen/canodb/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}

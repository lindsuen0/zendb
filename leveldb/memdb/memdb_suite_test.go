package memdb

import (
	"testing"

	"github.com/lindsuen0/zendb/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}

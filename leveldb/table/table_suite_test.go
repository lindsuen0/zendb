package table

import (
	"testing"

	"github.com/lindsuen0/zendb/leveldb/testutil"
)

func TestTable(t *testing.T) {
	testutil.RunSuite(t, "Table Suite")
}

package table

import (
	"testing"

	"github.com/lindsuen/canodb/leveldb/testutil"
)

func TestTable(t *testing.T) {
	testutil.RunSuite(t, "Table Suite")
}

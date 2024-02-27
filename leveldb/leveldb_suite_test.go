package leveldb

import (
	"testing"

	"github.com/lindsuen0/canodb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}

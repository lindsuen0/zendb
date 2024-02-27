package leveldb

import (
	"testing"

	"github.com/lindsuen/canodb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}

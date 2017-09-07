package local_test

import (
	"github.com/squioc/notebook/local"
	"github.com/squioc/notebook/testutil"
	"testing"
)

func TestLocalNotebook(t *testing.T) {
	// Arrange
	td, closer := testutil.TempDir(t)
	defer closer.Close()
	localNotebook := local.OpenNotebook(td)
	testutil.TestNotebook(localNotebook, t)
}

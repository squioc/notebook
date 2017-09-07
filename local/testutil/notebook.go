package testutil

import (
	"github.com/squioc/notebook"
	"github.com/squioc/notebook/local"
	"github.com/squioc/notebook/testutil"
	"io"
	"testing"
)

// TestLocalNotebook creates a temporary local notebook for tests
func TestLocalNotebook(t testing.TB) (notebook.Notebook, io.Closer) {
	td, closer := testutil.TempDir(t)
	return local.OpenNotebook(td), closer
}

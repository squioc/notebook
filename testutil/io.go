package testutil

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type closer struct {
	closeFunc func() error
}

// Close calls the callback
func (c *closer) Close() error {
	return c.closeFunc()
}

// NewCloser creates a new io.Closer to wrap a callback
func NewCloser(closeFunc func() error) io.Closer {
	return &closer{
		closeFunc: closeFunc,
	}
}

// TempDir creates a temporary for tests
func TempDir(t testing.TB) (string, io.Closer) {
	path, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("Cannot create temporary directory: %s", err)
	}

	return path, NewCloser(func() error { return os.RemoveAll(path) })
}

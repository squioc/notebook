package testutil

import (
	"io"
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

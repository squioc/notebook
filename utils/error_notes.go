package utils

import (
	"github.com/squioc/notebook"
)

// errorNote is a struct implementing the Notes interface
// and holding an error
type errorNotes struct {
	err error
}

// Next moves to the next note
func (n *errorNotes) Next() error {
	return n.err
}

// Name returns the name of the current note
func (n *errorNotes) Name() string {
	return ""
}

// Note returns the current note
func (n *errorNotes) Note() *notebook.Note {
	return nil
}

func NewErrorNotes(err error) notebook.Notes {
	return &errorNotes{
		err: err,
	}
}

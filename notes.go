package notebook

import (
	"errors"
)

var ErrorEndOfNotesReached = errors.New("The end of notes was reached")

// Notes is an interface holding a
// list of notes
type Notes interface {
	// Next moves to the next note
	Next() error

	// Name returns the name of the current note
	Name() string

	// Note returns the current note
	Note() *Note
}

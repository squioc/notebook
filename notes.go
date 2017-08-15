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

    // Value returns the current note with its name
    Value() (string, Note)
}

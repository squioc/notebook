package notebook

import (
    "errors"
)

var ErrorEndOfNotebooksReached = errors.New("The end of notebooks was reached")

// Notebooks is an interface holding a
// list of notebooks
type Notebooks interface {
    // Next moves to the nextbook note
    Next() error

    // Value returns the current notebook with its name
    Value() (string, Notebook)
}

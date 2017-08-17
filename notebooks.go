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

    // Name returns the name of the current notebook
    Name() string

    // Notebook returns the current notebook
    Notebook() Notebook
}

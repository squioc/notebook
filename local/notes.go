package local

import (
	"github.com/squioc/notebook"
	"os"
	"path/filepath"
)

// localNotes allows to iterate over a list of local notes
type localNotes struct {
	baseDirectory string

	namesOfNotes []string

	currentPosition int

	currentNote *notebook.Note
}

// loadCurrentNote reads the content of the current note
func (n *localNotes) loadCurrentNote() (*notebook.Note, error) {
	path := filepath.Join(n.baseDirectory, n.Name())
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return LoadNote(file)
}

// Next moves to the next note
func (n *localNotes) Next() error {
	if n.currentPosition+1 < len(n.namesOfNotes) {
		var err error
		n.currentPosition = n.currentPosition + 1
		n.currentNote, err = n.loadCurrentNote()
		return err
	}

	return notebook.ErrorEndOfNotesReached
}

// Name returns the name of the current note
func (n *localNotes) Name() string {
	return n.namesOfNotes[n.currentPosition]
}

// Note returns the content of the current note
func (n *localNotes) Note() *notebook.Note {
	return n.currentNote
}

// newLocalNotes creates a new iterator over local notes
func newLocalNotes(baseDir string, namesOfNotes []string) notebook.Notes {
	return &localNotes{
		baseDirectory:   baseDir,
		namesOfNotes:    namesOfNotes,
		currentPosition: 0,
	}
}

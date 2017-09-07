package local

import (
	"github.com/squioc/notebook"
	"github.com/squioc/notebook/utils"
	"os"
	"path/filepath"
)

// localNotebook is a struct for notebook
// on a filesystem
type localNotebook struct {
	path string
}

// Get returns a note from its name
func (n *localNotebook) Get(name string) (*notebook.Note, error) {
	path := filepath.Join(n.path, name)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return LoadNote(file)
}

// Set adds or updates a note in the notebook
func (n *localNotebook) Set(name string, note *notebook.Note) error {
	path := filepath.Join(n.path, name)

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	_, err = SaveNote(note, file)
	return err
}

// Delete deletes a note from the notebook
func (n *localNotebook) Delete(name string) error {
	path := filepath.Join(n.path, name)
	return os.Remove(path)
}

// List returns a list of name of notes holding in the notebook
func (n *localNotebook) List() notebook.Notes {
	var notes []string
	err := filepath.Walk(n.path, func(path string, info os.FileInfo, err error) error {
		_, filename := filepath.Split(path)
		notes = append(notes, filename)
		return nil
	})
	if err != nil {
		return utils.NewErrorNotes(err)
	}
	return newLocalNotes(n.path, notes)
}

// OpenNotebook opens a local notebook
func OpenNotebook(path string) notebook.Notebook {
	return &localNotebook{
		path: path,
	}
}

package testutil

import (
	"github.com/squioc/notebook"
	"testing"
)

func ExpectEmptyNotebook(nb notebook.Notebook, t testing.TB) {
	notes := nb.List()
	err := notes.Next()
	if err != notebook.ErrorEndOfNotesReached {
		t.Errorf("The notebook is not empty: %s", err)
	}
}

func ExpectNonEmptyNotebook(nb notebook.Notebook, t testing.TB) {
	notes := nb.List()
	err := notes.Next()
	if err != nil {
		t.Error("The notebook is empty")
	}
}

func TestNoteAdditionToNotebook(nb notebook.Notebook, name string, note *notebook.Note, t testing.TB) {
	err := nb.Set(name, note)
	if err != nil {
		t.Errorf("Cannot add the note in the notebook: %s", err)
	}
}

func TestNoteRetrievalFromNotebook(nb notebook.Notebook, name string, t testing.TB) {
	_, err := nb.Get(name)
	if err != nil {
		t.Errorf("Cannot get the note from the notebook: %s", err)
	}
}

func TestNoteDeletionFromNotebook(nb notebook.Notebook, name string, t testing.TB) {
	err := nb.Delete(name)
	if err != nil {
		t.Errorf("Cannot delete the note from the notebook: %s", err)
	}
}

func TestNotebook(nb notebook.Notebook, t testing.TB) {
	note := &notebook.Note{}

	// Act & Assert
	ExpectEmptyNotebook(nb, t)
	TestNoteAdditionToNotebook(nb, "test", note, t)
	TestNoteRetrievalFromNotebook(nb, "test", t)
	ExpectNonEmptyNotebook(nb, t)
	TestNoteDeletionFromNotebook(nb, "test", t)
	ExpectEmptyNotebook(nb, t)
}

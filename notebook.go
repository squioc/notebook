package notebook

// Notebook is an interface which represents
// a collection of notes
type Notebook interface {
    // List returns the list of notes holding in the notebook
    List() Notes

    // Get returns a note from the notebook according to its identifier
    Get([]byte) (Note, error)

    // Set adds or updates a note in the notebook
    Set([]byte, []byte) error

    // Delete deletes a note from the notebook according to its identifier
    Delete([]byte) error
}

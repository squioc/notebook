package notebook

// Notebook is an interface which represents
// a collection of notes
type Notebook interface {
    // List returns a list of notes holding in the notebook
    List() Notes

    // Get returns a note from the notebook according to its name
    Get(string) (*Note, error)

    // Set adds or updates a note in the notebook
    Set(string, *Note) error

    // Delete deletes a note from the notebook according to its name
    Delete(string) error
}

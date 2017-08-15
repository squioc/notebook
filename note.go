package notebook

// Note is a struct which represents
// a information to keep
type Note struct {
    // The category of the note
    Category string

    // The content of the note
    Content []byte
}

// NewNote creates a new note from a category and a content
func NewNote(Category string, Content []byte) *Note {
    return &Note {
        Category: Category,
        Content: Content,
    }
}

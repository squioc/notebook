package local

import (
    "bytes"
    "errors"
    "io"
    "io/ioutil"
    "github.com/squioc/notebook"
)

var ErrorInvalidLocalNoteFormat = errors.New("The format of the local note is invalid. Missing the NUL separator between the type of the note and its content")

// LoadNote reads a note from a io.Reader
func LoadNote(reader io.Reader) (*notebook.Note, error) {
    content, err := ioutil.ReadAll(reader)
    if err != nil {
        return nil, err
    }

    separatorPosition := bytes.IndexByte(content, 0x00)
    if separatorPosition == -1 {
        return nil, ErrorInvalidLocalNoteFormat
    }

    return notebook.NewNote(
        string(content[:separatorPosition]),
        content[separatorPosition+1:],
    ), nil
}

// SaveNote writes a note in a io.Writer
func SaveNote(note *notebook.Note, writer io.Writer) (int64, error) {
    var buffer bytes.Buffer
    buffer.WriteString(note.Category)
    buffer.WriteByte(0x0)
    buffer.Write(note.Content)
    return buffer.WriteTo(writer)
}


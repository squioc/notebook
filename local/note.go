package local

import (
	"bytes"
	"errors"
	"github.com/squioc/notebook"
	"io"
	"io/ioutil"
)

var ErrorInvalidLocalNoteFormat = errors.New("The format of the local note is invalid. Missing the NUL separator between the type of the note and its content")

// LoadNote reads a note from a io.Reader
func LoadNote(reader io.Reader) (*notebook.Note, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	separatorPosition := bytes.IndexByte(data, 0x00)
	if separatorPosition == -1 {
		return nil, ErrorInvalidLocalNoteFormat
	}

	category := string(data[:separatorPosition])
	content := data[separatorPosition+1:]
	return notebook.NewNote(category, content), nil
}

// SaveNote writes a note in a io.Writer
func SaveNote(note *notebook.Note, writer io.Writer) (int64, error) {
	var buffer bytes.Buffer
	buffer.WriteString(note.Category)
	buffer.WriteByte(0x0)
	buffer.Write(note.Content)
	return buffer.WriteTo(writer)
}

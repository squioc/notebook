package local_test

import (
    "bytes"
    "testing"
    "reflect"
    "github.com/squioc/notebook"
    "github.com/squioc/notebook/local"
)

func TestLoadNote(t *testing.T) {
    // Arrange
    data := []byte{'p', 'o', 's', 't', 0x0, 'c', 'o', 'n', 't', 'e', 'n', 't'}
    reader := bytes.NewBuffer(data)
    expectedCategory := "post"
    expectedContent := []byte("content")

    // Act
    note, err := local.LoadNote(reader)

    // Arrange
    if err != nil {
        t.Fatalf("An error occured while reading the note: %v", err)
    }
    actualCategory := note.Category
    if !reflect.DeepEqual(actualCategory, expectedCategory) {
        t.Errorf("Mismatching note category. Expected: '%s'. Got: '%s'", expectedCategory, actualCategory)
    }
    actualContent := note.Content
    if !reflect.DeepEqual(actualContent, expectedContent) {
        t.Errorf("Mismatching note content. Expected: '%s'. Got: '%s'", expectedContent, actualContent)
    }
}

func TestSaveNote(t *testing.T) {
    // Arrange
    noteCategory := "comment"
    content := []byte("content")
    note := notebook.NewNote(noteCategory, content)
    expectedData := []byte{'c', 'o', 'm', 'm', 'e', 'n', 't', 0x0, 'c', 'o', 'n', 't', 'e', 'n', 't'}
    var buffer bytes.Buffer

    // Act
    nbBytesWrited, err := local.SaveNote(note, &buffer)

    // Assert
    if err != nil {
        t.Fatalf("An error occured while writing the note: %v", err)
    }
    if nbBytesWrited == 0 {
        t.Errorf("No bytes writes")
    }
    actualData := buffer.Bytes()
    if !reflect.DeepEqual(actualData, expectedData) {
        t.Errorf("Mismatching data. Expected: '%v'. Got: '%v'", expectedData, actualData)
    }
}

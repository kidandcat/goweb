package models

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB_NAME = "file::memory:?cache=shared"

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestNoteRepository_First(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(DB_NAME))
	if err != nil {
		t.Fatal(err)
	}

	repo := NewNoteRepository(db)

	// Create a sample note
	note := &Note{Content: "Test Note"}
	repo.Save(note)

	// Call the First method
	result := repo.First()

	// Verify the result
	if result == nil {
		t.Error("Expected a note, but got nil")
	} else if result.Content != note.Content {
		t.Errorf("Expected content '%s', but got '%s'", note.Content, result.Content)
	}
}

func TestNoteRepository_Save(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(DB_NAME))
	if err != nil {
		t.Fatal(err)
	}
	repo := NewNoteRepository(db)
	// Create a sample note
	note := &Note{
		Content: "Test Note",
	}
	repo.Save(note)
	// Update the note's content
	id := note.ID
	newContent := "Updated Note"
	note.Content = newContent
	repo.Save(note)
	// Retrieve the note again
	result := repo.Find(id)
	// Verify the updated content
	if result == nil {
		t.Error("Expected a note, but got nil")
	} else if result.Content != newContent {
		t.Errorf("Expected content '%s', but got '%s'", newContent, result.Content)
	}
}

func TestNoteRepository_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(DB_NAME))
	if err != nil {
		t.Fatal(err)
	}
	repo := NewNoteRepository(db)
	// Create a sample note
	note := &Note{Content: "Test Note"}
	repo.Save(note)
	// Retrieve the note again
	result := repo.First()
	// Verify the created note
	if result == nil {
		t.Error("Expected a note, but got nil")
	} else if result.Content != note.Content {
		t.Errorf("Expected content '%s', but got '%s'", note.Content, result.Content)
	}
}

func TestNewNoteRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(DB_NAME))
	if err != nil {
		t.Fatal(err)
	}

	repo := NewNoteRepository(db)

	// Verify that the NoteRepository instance is created correctly
	if repo.DB != db {
		t.Error("Expected DB to be set correctly")
	}
}

func TestNoteRepository_Find(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(DB_NAME))
	if err != nil {
		t.Fatal(err)
	}
	repo := NewNoteRepository(db)

	// Create a test note
	expectedContent := "Test Note"
	note := &Note{Content: expectedContent}
	repo.Save(note)

	// Call the Find method
	foundNote := repo.Find(note.ID)

	// Check if the found note matches the expected note
	assert.NotNil(t, foundNote)
	assert.Equal(t, expectedContent, foundNote.Content)
}

func TestNoteRepository_FindNil(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(DB_NAME))
	if err != nil {
		t.Fatal(err)
	}
	repo := NewNoteRepository(db)

	// Call the Find method
	foundNote := repo.Find(0)
	assert.Nil(t, foundNote)
}

package services

import (
	"goweb/backend/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewNotesService(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"))
	if err != nil {
		t.Fatal(err)
	}

	service := NewNotesService(db)

	assert.NotNil(t, service)
	assert.NotNil(t, service.db)
	assert.NotNil(t, service.repository)
}

func TestNotesService_Read(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"))
	if err != nil {
		t.Fatal(err)
	}
	service := NewNotesService(db)

	// Create a test note
	expectedContent := "Test Note"
	note := &models.Note{Content: expectedContent}
	service.repository.Save(note)

	// Call the Read method
	content, err := service.Read()

	// Check if the returned content matches the expected content
	assert.NoError(t, err)
	assert.Equal(t, expectedContent, content)
}

func TestNotesService_Write(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"))
	if err != nil {
		t.Fatal(err)
	}
	service := NewNotesService(db)
	expectedContent := "Test Note"
	service.Write(expectedContent)

	// Retrieve the note from the database
	note := service.repository.First()
	assert.NotNil(t, note)
	assert.Equal(t, expectedContent, note.Content)

	// Update the note content
	updatedContent := "Updated Note"
	service.Write(updatedContent)

	// Retrieve the updated note from the database
	updatedNote := service.repository.First()
	assert.NotNil(t, updatedNote)
	assert.Equal(t, updatedContent, updatedNote.Content)
}

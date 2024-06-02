package services

import (
	"goweb/backend/models"

	"gorm.io/gorm"
)

var NotesClient struct {
	Read  func() (string, error)
	Write func(string) error
}

type NotesService struct {
	db         *gorm.DB
	repository models.NoteRepository
}

func NewNotesService(db *gorm.DB) *NotesService {
	return &NotesService{
		db:         db,
		repository: *models.NewNoteRepository(db),
	}
}

func (h *NotesService) Read() (string, error) {
	n := h.repository.First()
	return n.Content, nil
}

func (h *NotesService) Write(in string) {
	n := h.repository.First()
	if n == nil {
		h.repository.Create(&models.Note{Content: in})
	} else {
		n.Content = in
		h.repository.Save(n)
	}
}

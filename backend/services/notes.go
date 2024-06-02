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
	db *gorm.DB
}

func NewNotesService(db *gorm.DB) *NotesService {
	return &NotesService{db: db}
}

func (h *NotesService) Read() string {
	var note models.Note
	h.db.First(&note)
	return note.Content
}

func (h *NotesService) Write(in string) {
	var n models.Note
	res := h.db.First(&n)
	if res.RowsAffected == 0 {
		h.db.Create(&models.Note{Content: in})
	} else {
		n.Content = in
		h.db.Save(&n)
	}
}

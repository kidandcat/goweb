package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Content string
}

type NoteRepository struct {
	DB *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
	db.AutoMigrate(&Note{})
	return &NoteRepository{DB: db}
}

func (r *NoteRepository) First() *Note {
	var note Note
	err := r.DB.First(&note).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return &note
}

func (r *NoteRepository) Save(note *Note) {
	err := r.DB.Save(note).Error
	if err != nil {
		panic(err)
	}
}

func (r *NoteRepository) Find(id uint) *Note {
	var note Note
	err := r.DB.First(&note, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return &note
}

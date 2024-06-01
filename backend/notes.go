package backend

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Content string
}

var NotesClient struct {
	Read  func() (string, error)
	Write func(string) error
}

type NotesService struct{}

func (h *NotesService) Read() string {
	var note Note
	db.First(&note)
	return note.Content
}

func (h *NotesService) Write(in string) {
	var n Note
	res := db.First(&n)
	if res.RowsAffected == 0 {
		db.Create(&Note{Content: in})
	} else {
		n.Content = in
		db.Save(&n)
	}
}

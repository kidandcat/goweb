package backend

var NotesClient struct {
	Read  func() string
	Write func(string) error
}

type Notes struct {
	note string
}

func (h *Notes) Read() string {
	return h.note
}

func (h *Notes) Write(in string) {
	h.note = in
}

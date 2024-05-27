package backend

var APIClient struct {
	Read  func() string
	Write func(string) error
}

type API struct {
	note string
}

func (h *API) Read() string {
	return h.note
}

func (h *API) Write(in string) {
	h.note = in
}

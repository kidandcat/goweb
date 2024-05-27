package main

type API struct {
	note string
}

var api struct {
	Read  func() string
	Write func(string) error
}

func (h *API) Read() string {
	return h.note
}

func (h *API) Write(in string) {
	h.note = in
}

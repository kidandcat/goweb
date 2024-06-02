package frontend

import (
	"fmt"
	"goweb/backend/services"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Hello struct {
	app.Compo
	initialNote string
}

func (h *Hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Hello World!"),
		&Clock{},
		app.Textarea().
			Attr("value", h.initialNote).
			OnInput(func(ctx app.Context, e app.Event) {
				text := ctx.JSSrc().Get("value").String()
				err := services.NotesClient.Write(text)
				fmt.Println("api.Write", text, err)
			}),
	)
}

func (h *Hello) OnMount(ctx app.Context) {
	var err error
	h.initialNote, err = services.NotesClient.Read()
	if err != nil {
		fmt.Println("NotesClient.Read", err)
	}
}

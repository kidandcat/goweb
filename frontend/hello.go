package frontend

import (
	"fmt"
	"goweb/backend"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Hello struct {
	app.Compo
	initialNote string
}

func (h *Hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Hello World!"),
		app.Textarea().
			Attr("value", h.initialNote).
			OnInput(func(ctx app.Context, e app.Event) {
				text := ctx.JSSrc().Get("value").String()
				err := backend.APIClient.Write(text)
				fmt.Println("api.Write", text, err)
			}),
	)
}

func (h *Hello) OnMount(ctx app.Context) {
	h.initialNote = backend.APIClient.Read()
}

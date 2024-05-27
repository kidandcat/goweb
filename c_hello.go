package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type hello struct {
	app.Compo
	initialNote string
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Hello World!"),
		app.Textarea().
			Attr("value", h.initialNote).
			OnInput(func(ctx app.Context, e app.Event) {
				text := ctx.JSSrc().Get("value").String()
				err := api.Write(text)
				fmt.Println("api.Write", text, err)
			}),
	)
}

func (h *hello) OnMount(ctx app.Context) {
	h.initialNote = api.Read()
}

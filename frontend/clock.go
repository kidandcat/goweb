package frontend

import (
	"fmt"
	"goweb/backend/services"
	"time"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Clock struct {
	app.Compo
	now time.Time
}

func (h *Clock) Render() app.UI {
	return app.H2().
		Class("clock").
		Text(fmt.Sprintf("Clock: %v", h.now.Format(time.RFC3339)))
}

func (h *Clock) OnMount(ctx app.Context) {
	h.Update(ctx)
	ctx.Async(func() {
		for range time.Tick(time.Second) {
			ctx.Dispatch(func(ctx app.Context) {
				h.Update(ctx)
			})
		}
	})
}

func (h *Clock) Update(ctx app.Context) {
	now, err := services.ClockClient.Read()
	if err != nil {
		fmt.Println("api.Now", err)
	}
	h.now = now
}

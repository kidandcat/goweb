package frontend

import "github.com/maxence-charriere/go-app/v10/pkg/app"

func RegisterRoutes() {
	app.Route("/hello", app.NewZeroComponentFactory(&Hello{}))
}

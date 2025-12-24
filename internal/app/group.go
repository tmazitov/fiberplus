package app

import (
	"github.com/gofiber/fiber/v2"
)

type Group[Services any] struct {
	core fiber.Router
	app  *App[Services]
}

func NewGroup[Services any](app *App[Services], route string, handlers ...fiber.Handler) *Group[Services] {
	return &Group[Services]{
		core: app.core.Group(route, handlers...),
		app:  app,
	}
}

func (g *Group[Services]) Add(e *Endpoint[Services]) *Group[Services] {

	e.Handler.Init()

	intro := e.Handler.IntroMods()
	main := e.Handler.Handle(g.app)
	outro := e.Handler.OutroMods()

	var pipeline = make([]fiber.Handler, 0, len(intro)+len(outro)+1)

	pipeline = append(pipeline, intro...)
	pipeline = append(pipeline, main)
	pipeline = append(pipeline, outro...)

	g.core.Add(e.Method, e.Route, pipeline...)

	return g
}

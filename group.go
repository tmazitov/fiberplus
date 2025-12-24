package fiberplus

import (
	"github.com/gofiber/fiber/v2"
)

type Group[Services any] struct {
	core fiber.Router
	app  *App[Services]
}

func newGroup[Services any](app *App[Services], route string, handlers ...fiber.Handler) *Group[Services] {
	return &Group[Services]{
		core: app.core.Group(route, handlers...),
		app:  app,
	}
}

func newSubGroup[Services any](parent *Group[Services], route string, handlers ...fiber.Handler) *Group[Services] {
	return &Group[Services]{
		app:  parent.app,
		core: parent.core.Group(route, handlers...),
	}
}

func (g *Group[Services]) Group(route string, handlers ...fiber.Handler) *Group[Services] {
	return newSubGroup(g, route, handlers...)
}

func (g *Group[Services]) Add(e *Endpoint[Services]) *Group[Services] {

	e.Handler.Init()

	mods := e.Handler.Mods()
	main := e.Handler.Handle(g.app)

	var pipeline = make([]fiber.Handler, 0, len(mods)+1)

	pipeline = append(pipeline, mods...)
	pipeline = append(pipeline, main)

	g.core.Add(e.Method, e.Route, pipeline...)

	return g
}

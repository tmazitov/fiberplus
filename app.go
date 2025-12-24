package fiberplus

import (
	"github.com/gofiber/fiber/v2"
)

type App[Services any] struct {
	core     *fiber.App
	services *Services
}

type AppConfig[Services any] struct {
	Core     fiber.Config
	Services *Services
}

func NewApp[Services any](config *AppConfig[Services]) *App[Services] {
	return &App[Services]{
		core:     fiber.New(config.Core),
		services: config.Services,
	}
}

func (a *App[Services]) Services() *Services {
	return a.services
}

func (a *App[Services]) Core() *fiber.App {
	return a.core
}

func (a *App[Services]) Group(route string, handlers ...fiber.Handler) *Group[Services] {
	return newGroup(a, route, handlers...)
}

func (a *App[Services]) Add(e *Endpoint[Services]) *App[Services] {

	e.Handler.Init()

	mods := e.Handler.Mods()
	main := e.Handler.Handle(a)

	var pipeline = make([]fiber.Handler, 0, len(mods)+1)

	pipeline = append(pipeline, mods...)
	pipeline = append(pipeline, main)

	a.core.Add(e.Method, e.Route, pipeline...)

	return a
}

func (a *App[Services]) Run(addr string) error {
	return a.core.Listen(addr)
}

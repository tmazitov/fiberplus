package app

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
	return NewGroup(a, route, handlers...)
}

func (a *App[Services]) Add(e *Endpoint[Services]) {
	e.Handler.Init()

	intro := e.Handler.IntroMods()
	main := e.Handler.Handle(a)
	outro := e.Handler.OutroMods()

	var pipeline = make([]fiber.Handler, 0, len(intro)+len(outro)+1)

	pipeline = append(pipeline, intro...)
	pipeline = append(pipeline, main)
	pipeline = append(pipeline, outro...)

	a.core.Add(e.Method, e.Route, pipeline...)
}

func (a *App[Services]) Run(addr string) error {
	return a.core.Listen(addr)
}

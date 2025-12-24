package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/tmazitov/fiberplus/internal/app"
	"github.com/tmazitov/fiberplus/internal/behavior"
)

type Database struct{}

func (d *Database) Test() string {
	return "Hello world!"
}

type Services struct {
	Database Database
}

type ReadHandlerRequest struct {
	Field string `json:"field" validate:"required,min=5,max=20"`
}

type ReadHandlerExample struct {
	behavior.ReadHandler[Services, ReadHandlerRequest]
}

type ReadHandlerResponse struct {
	Request *ReadHandlerRequest `json:"request"`
	Message string              `json:"message"`
}

func (h *ReadHandlerExample) Handle(app *app.App[Services]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		log.Info("input says:")

		input := ctx.Locals("Input").(*ReadHandlerRequest)

		return ctx.JSON(&ReadHandlerResponse{Request: input, Message: "ok!"})
	}
}

func main() {
	var (
		services = &Services{
			Database: Database{},
		}
		a = app.NewApp(&app.AppConfig[Services]{
			Services: services,
		})
	)
	a.Core().Use(logger.New())

	a.Group("/test").
		Add(&app.Endpoint[Services]{Method: "POST", Route: "/in", Handler: &ReadHandlerExample{}})

	log.Errorf("app error : ", a.Run(":8070"))
}

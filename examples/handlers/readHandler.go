package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/tmazitov/fiberplus/examples/services"
	fiberplus "github.com/tmazitov/fiberplus/internal"
	"github.com/tmazitov/fiberplus/internal/behavior"
)

type ReadHandlerRequest struct {
	Field string `json:"field" validate:"required,min=5,max=20"`
}

type ReadHandlerExample struct {
	behavior.ReadHandler[services.Services, ReadHandlerRequest]
}

type ReadHandlerResponse struct {
	Request *ReadHandlerRequest `json:"request"`
	Message string              `json:"message"`
}

func (h *ReadHandlerExample) Handle(app *fiberplus.App[services.Services]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		log.Info("input says:")

		input := ctx.Locals("Input").(*ReadHandlerRequest)

		return ctx.JSON(&ReadHandlerResponse{Request: input, Message: "ok!"})
	}
}

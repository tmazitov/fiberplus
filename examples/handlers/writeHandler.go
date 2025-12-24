package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tmazitov/fiberplus/examples/services"
	fiberplus "github.com/tmazitov/fiberplus/internal"
	"github.com/tmazitov/fiberplus/internal/behavior"
)

type WriteHandlerResponse struct {
	Message string `json:"message"`
}

type WriteHandlerExample struct {
	behavior.WriteHandler[services.Services, *WriteHandlerResponse]
}

func (h *WriteHandlerExample) Handle(app *fiberplus.App[services.Services]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals("Output", &WriteHandlerResponse{Message: "cow say: moo!"})
		return nil
	}
}

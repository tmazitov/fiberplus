package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tmazitov/fiberplus/examples/services"
	fiberplus "github.com/tmazitov/fiberplus/internal"
	"github.com/tmazitov/fiberplus/internal/behavior"
)

type CoreHandlerRequest struct {
	Message string `json:"message" validate:"required"`
}

type CoreHandlerResponse struct {
	Message string `json:"message"`
}

type CoreHandlerExample struct {
	behavior.CoreHandler[services.Services, *CoreHandlerRequest, *CoreHandlerResponse]
}

func (h *CoreHandlerExample) Handle(app *fiberplus.App[services.Services]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var (
			input *CoreHandlerRequest
			ok    bool
		)

		if input, ok = h.RequestBody(ctx); !ok {
			h.Reply(ctx, &CoreHandlerResponse{Message: "nah!"})
			return nil
		}

		h.Reply(ctx, &CoreHandlerResponse{Message: input.Message})

		return nil
	}
}

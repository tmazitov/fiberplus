package mods

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tmazitov/fiberplus/internal/validator"
)

type JsonInputMod[I any] struct{}

func (i *JsonInputMod[I]) Setup() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var (
			inputPayload I
		)
		if err := ctx.BodyParser(&inputPayload); err != nil {
			return fiber.ErrBadRequest
		}

		if err := validator.Instance().Struct(inputPayload); err != nil {
			return fiber.ErrBadRequest
		}
		ctx.Locals("Input", &inputPayload)
		return ctx.Next()
	}
}

func (i *JsonInputMod[O]) Type() ModeType {
	return IntroType
}

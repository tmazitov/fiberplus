package mods

import (
	"github.com/gofiber/fiber/v2"
)

type JsonOutputMod[O any] struct {
}

func (o *JsonOutputMod[O]) Setup() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if err := ctx.Next(); err != nil {
			return err
		}

		var (
			outputRaw     any = ctx.Locals("Output")
			outputPayload O
			ok            bool
		)

		if outputRaw == nil {
			return nil
		}

		if outputPayload, ok = outputRaw.(O); !ok {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to typecast output")
		}

		if err := ctx.JSON(outputPayload); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to convert the output to JSON format")
		}
		return nil
	}
}

func (o *JsonOutputMod[O]) Type() ModeType {
	return OutroType
}

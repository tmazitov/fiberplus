package mods

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type JsonOutputMod[O any] struct {
}

func (o *JsonOutputMod[O]) Setup() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var (
			outputRaw     any = ctx.Locals("Output", nil)
			outputPayload O
			ok            bool
		)

		if outputRaw == nil {
			return ctx.Next()
		}

		if outputPayload, ok = outputRaw.(O); !ok {
			return errors.New("jsonOutputMode error : failed to typecast output value")
		}

		if err := ctx.JSON(outputPayload); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return fmt.Errorf("jsonOutputMode error : failed to convert the output to JSON format : %w", err)
		}
		return ctx.Next()
	}
}

func (o *JsonOutputMod[O]) Type() ModeType {
	return OutroType
}

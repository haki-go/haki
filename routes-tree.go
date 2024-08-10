package haki

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (h *Haki) applyRoutesTree() {
	for path, route := range h.routes {
		for method, handlers := range route {
			var fiberHandlers []func(ctx *fiber.Ctx) error

			for _, handler := range handlers {
				fiberHandlers = append(fiberHandlers, applyHakiHandler(handler))
			}

			h.fiberApp.Add(method, path, fiberHandlers...)
		}
	}
}

func applyHakiHandler(handler RouteHandler) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		result := handler(Context{Request: ctx})

		if result != nil {
			switch v := result.(type) {
			case string:
				return ctx.SendString(v)
			case map[string]interface{}:
				return ctx.JSON(v)
			case Map:
				return ctx.JSON(v)
			case Array:
				return ctx.JSON(v)
			case Exception:
				return ctx.Status(v.StatusCode).SendString(v.Message)
			}

			fmt.Printf("Chegou aqui! Com Result: %v\n", result)

			if err := ctx.JSON(result); err != nil {
				fmt.Printf("Error: %v\n", err)
				return ctx.Status(500).SendString(err.Error())
			} else {
				ctx.Status(200).JSON(result)
			}
		}

		return ctx.Next()
	}
}

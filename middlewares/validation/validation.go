package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/haki-go/haki"
)

func ValidateBody[T any](dto T) func(c haki.Context) any {
	validate := validator.New()

	return func(c haki.Context) any {
		var body T

		if err := c.Request.BodyParser(&body); err != nil {
			return haki.Exception{Message: err.Error(), StatusCode: 400}
		}

		if err := validate.Struct(body); err != nil {
			return haki.Exception{Message: err.Error(), StatusCode: 400}
		}

		return nil
	}
}

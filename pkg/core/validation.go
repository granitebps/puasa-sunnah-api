package core

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"
)

type AppValidator struct {
	Validator *validator.Validate
}

type ErrorField struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func SetupValidator() *AppValidator {
	v := validator.New()

	// You can add your custom validator in here
	// v.RegisterValidation()

	return &AppValidator{
		Validator: v,
	}
}

func (v *AppValidator) Validate(c *fiber.Ctx, payload any) (fields []ErrorField, err error) {
	err = c.BodyParser(payload)
	if err != nil {
		return
	}

	err = v.Validator.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var elem ErrorField
			field := utils.CamelToSnake(err.Field())
			elem.Field = field
			elem.Message = validationRuleMessage(field, err.Tag(), err.Error(), err.Param())

			fields = append(fields, elem)
		}

		err = errors.New(fields[0].Message)
		return
	}

	return
}

func validationRuleMessage(field, rule, def, param string) string {
	switch rule {
	case "required":
		return fmt.Sprintf("The %s field is required.", field)
	case "email":
		return fmt.Sprintf("The %s field must be a valid email address.", field)
	case "min":
		return fmt.Sprintf("The %s field must be at least %s characters.", field, param)
	case "url":
		return fmt.Sprintf("The %s field must be a valid URL.", field)
	default:
		return def
	}
}

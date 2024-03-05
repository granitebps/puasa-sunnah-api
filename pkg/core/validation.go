package core

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/ansel1/merry/v2"
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
	v.RegisterValidation("isYear", validateYear)

	return &AppValidator{
		Validator: v,
	}
}

func (v *AppValidator) Validate(c *fiber.Ctx, payload any) (fields []ErrorField, err error) {
	err = c.BodyParser(payload)
	if err != nil {
		err = merry.Wrap(err)
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
	case "len":
		return fmt.Sprintf("The %s field must be at exactly %s characters long.", field, param)
	case "url":
		return fmt.Sprintf("The %s field must be a valid URL.", field)
	case "min":
		return fmt.Sprintf("The %s field must be at least %s.", field, param)
	case "max":
		return fmt.Sprintf("The %s may not be greater than %s.", field, param)
	case "datetime":
		if param != "" {
			return fmt.Sprintf("The %s field does not match the format %s.", field, param)
		}
		return fmt.Sprintf("The %s field must be a valid date/datetime", field)
	case "isYear":
		return fmt.Sprintf("The %s field must be a valid year (YYYY) and greater than 1900.", field)
	default:
		return def
	}
}

func validateYear(fl validator.FieldLevel) bool {
	yearType := fl.Field().Type().String()
	var yearStr string
	if yearType != "string" {
		yearStr = strconv.Itoa(int(fl.Field().Uint()))
	} else {
		yearStr = fl.Field().String()
	}

	// Regular expression to validate format (YYYY)
	regexPattern := `^[0-9]{4}$`
	if !regexp.MustCompile(regexPattern).MatchString(yearStr) {
		return false
	}

	// Your custom validation logic here (e.g., check for valid year range, leap year)
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return false
	}
	minYear := 1900
	maxYear := time.Now().Year() // Update with your desired max year
	if year < minYear || year > maxYear {
		return false
	}

	return true
}

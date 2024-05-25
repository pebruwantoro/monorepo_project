package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func SetupValidator() *validator.Validate {
	v := validator.New()

	// Register your custom validator here
	// https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Custom_Validation_Functions
	// v.RegisterValidation("name of your custom tag", validationFunc)

	return v
}

func Validate(c echo.Context, s interface{}) (err error) {
	if err = c.Bind(s); err != nil {
		err = fmt.Errorf("error bind : %s", err.Error())
		return
	}

	if err = c.Validate(s); err != nil {
		err = fmt.Errorf("error validate : %s", err.Error())
		return
	}

	return
}

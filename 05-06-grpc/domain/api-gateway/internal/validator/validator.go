package validator

import (
	"fmt"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func New() *Validator {
	v := &Validator{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}

	v.validator.RegisterValidation("timekitchen", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()

		_, err := time.Parse(time.Kitchen, value)
		return err == nil
	})

	// v.validator.RegisterValidation("date", func(fl validator.FieldLevel) bool {
	// 	value := fl.Field().Interface().(entity.Date)
	// 	// _, err := time.Parse(entity.Dateformat, value)
	// 	// fmt.Println(value)
	// 	// fmt.Println("AA", err)
	// 	// return err == nil
	// 	// value
	// })
	return v
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		errv, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}

		if len(errv) > 0 {
			err := errv[0]
			switch err.Tag() {
			case "required":
				return fmt.Errorf("%v: is required", err.Field())
			case "oneof":
				return fmt.Errorf("%v: must be one of %v", err.Field(), err.Param())
			case "min", "max":
				switch err.Kind() {
				case reflect.String:
					return fmt.Errorf("%v: must be %v %v char long", err.Field(), err.Tag(), err.Param())
				case reflect.Int:
					return fmt.Errorf("%v: must be %v %v", err.Field(), err.Tag(), err.Param())
				case reflect.Float64:
					return fmt.Errorf("%v: must be %v %v", err.Field(), err.Tag(), err.Param())
				default:
					return err
				}
			default:
				return err
			}
		}
	}
	return nil
}

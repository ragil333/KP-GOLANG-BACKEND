package helpers

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator"
)

func Validation(model interface{}) map[string]string {
	v := validator.New()
	errorList := make(map[string]string)
	if err := v.Struct(model); err != nil {
		err := err.(validator.ValidationErrors)
		for _, e := range err {
			var errMsg string
			field, _ := reflect.TypeOf(model).FieldByName(e.StructField())
			fieldName := field.Tag.Get("json")
			switch e.Tag() {
			case "required":
				errMsg = fmt.Sprintf("%s is required", fieldName)
			case "email":
				errMsg = fmt.Sprintf("%s is invalid", fieldName)
			case "min":
				errMsg = fmt.Sprintf("%s is invalid min %s", fieldName, e.Param())
			}
			errorList[fieldName] = errMsg
		}
		return errorList
	}
	return nil
}

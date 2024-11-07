package base

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type Api struct{}

func NewApi() *Api {
	return &Api{}
}

// ReadBody 读取body并校验
// desc: 待重构
func (*Api) ReadBody(ctx iris.Context, ptr any) error {

	if err := ctx.ReadBody(ptr); err != nil {
		invalid, ok := err.(*validator.InvalidValidationError)
		if ok {
			return invalid
		}
		validationErrs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, validationErr := range validationErrs {
				fieldName, t := validationErr.Field(), reflect.TypeOf(ptr)
				if t.Kind() == reflect.Ptr {
					t = t.Elem()
				}
				field, ok := t.FieldByName(fieldName) //通过反射获取filed
				if ok {
					errorInfo := field.Tag.Get("error") // 获取field对应的error tag值
					if errorInfo == "" {
						return validationErr
					}
					return errors.New(errorInfo) //返回错误
				} else {
					return validationErr
				}
			}
		}
		return err
	}
	return nil
}

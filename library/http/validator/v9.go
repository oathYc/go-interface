package validator

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &DefaultValidator{}

func (v *DefaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyInit()

		if err := v.validate.Struct(obj); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				// 暂时只返回第一个，todo 使用 errors: [key: message, ...] 的形式返回多个错误
				// todo 验证错误 message 使用自定义语言包
				return errors.New(fmt.Sprintf(
					"validation  failed on '%s'",
					err.Tag()+" "+err.Param(),
				))
			}
			return err
		}
	}

	return nil
}

func (v *DefaultValidator) Engine() interface{} {
	v.lazyInit()
	return v.validate
}

func (v *DefaultValidator) lazyInit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		// add any custom validations etc. here
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

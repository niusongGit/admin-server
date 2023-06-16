package validator

import (
	"github.com/gookit/validate"
	"github.com/gookit/validate/locales/zhcn"
)

func Validate(data interface{}) error {
	v := validate.New(data)

	// only for current Validation
	zhcn.Register(v)

	if !v.Validate() {
		return v.Errors.OneError()
	}

	return nil
}

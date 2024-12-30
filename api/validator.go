package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/nishijun/simplebank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check if currency is supported
		return util.IsCurrencySupported(currency)
	}
	return false
}

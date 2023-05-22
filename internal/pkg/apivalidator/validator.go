package apiValidator

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

// DefaultValidator 生成一个自带中文翻译的validator
func DefaultValidator() *validator.Validate {
	validator := validator.New()
	zh := zh.New()
	uni := ut.New(zh)
	trans, _ := uni.GetTranslator("zh")
	_ = zhTrans.RegisterDefaultTranslations(validator, trans)
	return validator
}

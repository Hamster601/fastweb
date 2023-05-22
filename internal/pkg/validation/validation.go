package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"

	"github.com/Hamster601/fastweb/configs"

	"github.com/gin-gonic/gin/binding"
	enTranslation "github.com/go-playground/apivalidator/v10/translations/en"
	zhTranslation "github.com/go-playground/apivalidator/v10/translations/zh"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
)

var trans ut.Translator

func init() {
	lang := configs.Get().Language.Local

	if lang == configs.ZhCN {
		trans, _ = ut.New(zh.New()).GetTranslator("zh")
		if err := zhTranslation.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans); err != nil {
			fmt.Println("apivalidator zh translation error", err)
		}
	}

	if lang == configs.EnUS {
		trans, _ = ut.New(en.New()).GetTranslator("en")
		if err := enTranslation.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans); err != nil {
			fmt.Println("apivalidator en translation error", err)
		}
	}
}

func Error(err error) (message string) {
	if validationErrors, ok := err.(validator.ValidationErrors); !ok {
		return err.Error()
	} else {
		for _, e := range validationErrors {
			message += e.Translate(trans) + ";"
		}
	}
	return message
}

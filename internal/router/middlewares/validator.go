package middlewares

import (
	apiValidator "github.com/Hamster601/fastweb/internal/pkg/apivalidator"
	"github.com/gin-gonic/gin"
)

// 使用例子：https://juejin.cn/post/6847902214279659533

func Validator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiValidator := apiValidator.DefaultValidator()
		ctx.Set("apiParam", apiValidator)
		ctx.Next()
	}
}

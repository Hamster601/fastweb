package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

func OpentracingJager() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		span, tracectx := opentracing.StartSpanFromContext(ctx.Request.Context(), "hello")
		defer span.Finish()
		ctx.Set("trace", tracectx)
		ctx.Next()
	}
}

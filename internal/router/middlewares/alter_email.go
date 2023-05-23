package middlewares

import (
	"fmt"
	"github.com/Hamster601/fastweb/internal/code"
	"github.com/Hamster601/fastweb/internal/pkg/businesserror"
	"github.com/Hamster601/fastweb/pkg/pkglog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"runtime/debug"
)

func AlterEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stackInfo := string(debug.Stack())
				pkglog.ProjectLogger.Error("got panic", zap.String("panic", fmt.Sprintf("%+v", err)), zap.String("stack", stackInfo))
				ctx.AbortWithError(http.StatusBadRequest, businesserror.Error(
					http.StatusInternalServerError,
					code.ServerError,
					code.Text(code.ServerError)),
				)
			}

			// region 发生错误，进行返回
			if ctx.IsAborted() {
			}

			pkglog.ProjectLogger.Info("trace-log",
				zap.Any("method", ctx.Request.Method),
				zap.Any("path", ctx.Request.URL),
				zap.Any("http_code", ctx.Writer.Status()),
				zap.Any("business_code", 0),
				zap.Any("success", 1),
				zap.Any("cost_seconds", 1),
				zap.Any("trace_id", 1),
				zap.Any("trace_info", 1),
				zap.Error(nil))

			// endregion
		}()
		ctx.Next()
	}
}

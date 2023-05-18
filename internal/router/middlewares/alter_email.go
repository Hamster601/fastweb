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

				//if notifyHandler := opt.alertNotify; notifyHandler != nil {
				//	notifyHandler(&proposal.AlertMessage{
				//		ProjectName:  configs.ProjectName,
				//		Env:          env.Active().Value(),
				//		TraceID:      traceId,
				//		HOST:         context.Host(),
				//		URI:          context.URI(),
				//		Method:       context.Method(),
				//		ErrorMessage: err,
				//		ErrorStack:   stackInfo,
				//		Timestamp:    time.Now(),
				//	})
				//}
			}
			// endregion

			// region 发生错误，进行返回
			if ctx.IsAborted() {
				//for i := range ctx.Errors {
				//	multierr.AppendInto(&abortErr, ctx.Errors[i])
				//}
				//
				//if err := context.abortError(); err != nil { // customer err
				//	// 判断是否需要发送告警通知
				//	if err.IsAlert() {
				//		if notifyHandler := opt.alertNotify; notifyHandler != nil {
				//			notifyHandler(&proposal.AlertMessage{
				//				ProjectName:  configs.ProjectName,
				//				Env:          env.Active().Value(),
				//				TraceID:      traceId,
				//				HOST:         context.Host(),
				//				URI:          context.URI(),
				//				Method:       context.Method(),
				//				ErrorMessage: err.Message(),
				//				ErrorStack:   fmt.Sprintf("%+v", err.StackError()),
				//				Timestamp:    time.Now(),
				//			})
				//		}
				//	}
				//
				//	multierr.AppendInto(&abortErr, err.StackError())
				//	businessCode = err.BusinessCode()
				//	businessCodeMsg = err.Message()
				//	response = &code.Failure{
				//		Code:    businessCode,
				//		Message: businessCodeMsg,
				//	}
				//	ctx.JSON(err.HTTPCode(), response)
				//}
			}
			// endregion

			//logger.Info("trace-log",
			//	zap.Any("method", ctx.Request.Method),
			//	zap.Any("path", decodedURL),
			//	zap.Any("http_code", ctx.Writer.Status()),
			//	zap.Any("business_code", businessCode),
			//	zap.Any("success", t.Success),
			//	zap.Any("cost_seconds", t.CostSeconds),
			//	zap.Any("trace_id", t.Identifier),
			//	zap.Any("trace_info", t),
			//	zap.Error(abortErr),)

			// endregion
		}()
		ctx.Next()
	}
}

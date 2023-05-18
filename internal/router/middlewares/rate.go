package middlewares

import (
	"github.com/Hamster601/fastweb/configs"
	"github.com/Hamster601/fastweb/internal/code"
	"github.com/Hamster601/fastweb/internal/pkg/businesserror"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

var limiter = rate.NewLimiter(rate.Every(time.Second*1), configs.MaxRequestsPerSecond) // 计时器限流，每秒QPS10000

// LimitRate count limit middleware
func LimitRate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !limiter.Allow() {
			ctx.AbortWithError(http.StatusForbidden, businesserror.Error(
				http.StatusTooManyRequests,
				code.TooManyRequests,
				code.Text(code.TooManyRequests)),
			)
			return
		}
		ctx.Next()
	}
}

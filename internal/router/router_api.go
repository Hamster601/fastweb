package router

import (
	"github.com/Hamster601/fastweb/internal/api/admin"
	"github.com/Hamster601/fastweb/pkg/pkglog"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/Hamster601/fastweb/pkg/env"
)

func SystemHealth(ctx *gin.Context) {
	resp := &struct {
		Timestamp   time.Time `json:"timestamp"`
		Environment string    `json:"environment"`
		Host        string    `json:"host"`
		Status      string    `json:"status"`
	}{
		Timestamp:   time.Now(),
		Environment: env.Active().Value(),
		Host:        "ctx.",
		Status:      "ok",
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}

func initAPIRouter(eg *gin.Engine) {
	eg.Group("/system").GET("/health", SystemHealth)
	v1api := eg.Group("/v1/api")
	adminHandler := admin.New(pkglog.ProjectLogger)
	v1api.POST("/login", adminHandler.LoginNew)

}

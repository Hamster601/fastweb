package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/Hamster601/fastweb/internal/api/admin"
	"github.com/Hamster601/fastweb/internal/api/authorized"
	"github.com/Hamster601/fastweb/internal/api/config"
	"github.com/Hamster601/fastweb/internal/api/cron"
	"github.com/Hamster601/fastweb/internal/api/helper"
	"github.com/Hamster601/fastweb/internal/api/tool"
	"github.com/Hamster601/fastweb/internal/pkg/core"
	"github.com/Hamster601/fastweb/pkg/env"
)

func initAPIRouter(eg *gin.Engine) {
	eg.Group("/system").GET("/health", func(ctx *gin.Context) {
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
	})
	eg.Group("/v1/api")
}
func setApiRouter(r *resource) {
	// helper
	helperHandler := helper.New(r.logger, r.db, r.cache)

	helpers := r.mux.Group("/helper")
	{
		helpers.GET("/md5/:str", helperHandler.Md5())
		helpers.POST("/sign", helperHandler.Sign())
	}
	// 系统健康检查
	r.mux.Group("/system").GET("/health", func(ctx core.Context) {
		resp := &struct {
			Timestamp   time.Time `json:"timestamp"`
			Environment string    `json:"environment"`
			Host        string    `json:"host"`
			Status      string    `json:"status"`
		}{
			Timestamp:   time.Now(),
			Environment: env.Active().Value(),
			Host:        ctx.Host(),
			Status:      "ok",
		}
		ctx.Payload(resp)
	})
	api := r.mux.Group("/v1/api")
	// admin
	adminHandler := admin.New(r.logger, r.db, r.cache)

	// 需要签名验证，无需登录验证，无需 RBAC 权限验证
	r.mux.Group("/api", r.interceptors.CheckSignature()).POST("/login", adminHandler.Login())
	api.POST("/login", r.interceptors.CheckSignature(), adminHandler.Login())
	admin := api.Group("/admin", core.WrapAuthHandler(r.interceptors.CheckLogin), r.interceptors.CheckSignature())
	// 需要签名验证、登录验证，无需 RBAC 权限验证
	admin.POST("/admin/logout", adminHandler.Logout())
	admin.PATCH("/admin/modify_password", adminHandler.ModifyPassword())
	admin.GET("/admin/info", adminHandler.Detail())
	admin.PATCH("/admin/modify_personal_info", adminHandler.ModifyPersonalInfo())

	// 需要签名验证、登录验证、RBAC 权限验证
	rbac := api.Group("/rbac", core.WrapAuthHandler(r.interceptors.CheckLogin), r.interceptors.CheckSignature(), r.interceptors.CheckRBAC())
	authorizedHandler := authorized.New(r.logger, r.db, r.cache)
	rbac.POST("/authorized", authorizedHandler.Create())
	rbac.GET("/authorized", authorizedHandler.List())
	rbac.PATCH("/authorized/used", authorizedHandler.UpdateUsed())
	rbac.DELETE("/authorized/:id", core.AliasForRecordMetrics("/api/authorized/info"), authorizedHandler.Delete())

	rbac.POST("/authorized_api", authorizedHandler.CreateAPI())
	rbac.GET("/authorized_api", authorizedHandler.ListAPI())
	rbac.DELETE("/authorized_api/:id", core.AliasForRecordMetrics("/api/authorized_api/info"), authorizedHandler.DeleteAPI())

	rbac.POST("/admin", adminHandler.Create())
	rbac.GET("/admin", adminHandler.List())
	rbac.PATCH("/admin/used", adminHandler.UpdateUsed())
	rbac.PATCH("/admin/offline", adminHandler.Offline())
	rbac.PATCH("/admin/reset_password/:id", core.AliasForRecordMetrics("/api/admin/reset_password"), adminHandler.ResetPassword())
	rbac.DELETE("/admin/:id", core.AliasForRecordMetrics("/api/admin"), adminHandler.Delete())

	// tool
	toolHandler := tool.New(r.logger, r.db, r.cache)
	rbac.GET("/tool/hashids/encode/:id", core.AliasForRecordMetrics("/api/tool/hashids/encode"), toolHandler.HashIdsEncode())
	rbac.GET("/tool/hashids/decode/:id", core.AliasForRecordMetrics("/api/tool/hashids/decode"), toolHandler.HashIdsDecode())
	rbac.POST("/tool/cache/search", toolHandler.SearchCache())
	rbac.PATCH("/tool/cache/clear", toolHandler.ClearCache())
	rbac.GET("/tool/data/dbs", toolHandler.Dbs())
	rbac.POST("/tool/data/tables", toolHandler.Tables())
	rbac.POST("/tool/data/admin", toolHandler.SearchMySQL())
	rbac.POST("/tool/send_message", toolHandler.SendMessage())

	// config
	configHandler := config.New(r.logger, r.db, r.cache)
	rbac.PATCH("/config/email", configHandler.Email())

	// cron
	cronHandler := cron.New(r.logger, r.db, r.cache, r.cronServer)
	rbac.POST("/cron", cronHandler.Create())
	rbac.GET("/cron", cronHandler.List())
	rbac.GET("/cron/:id", core.AliasForRecordMetrics("/api/cron/detail"), cronHandler.Detail())
	rbac.POST("/cron/:id", core.AliasForRecordMetrics("/api/cron/modify"), cronHandler.Modify())
	rbac.PATCH("/cron/used", cronHandler.UpdateUsed())
	rbac.PATCH("/cron/exec/:id", core.AliasForRecordMetrics("/api/cron/exec"), cronHandler.Execute())

}

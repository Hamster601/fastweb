package router

import (
	"time"

	"github.com/Hamster601/fastweb/internal/api/admin"
	"github.com/Hamster601/fastweb/internal/api/authorized"
	"github.com/Hamster601/fastweb/internal/api/config"
	"github.com/Hamster601/fastweb/internal/api/cron"
	"github.com/Hamster601/fastweb/internal/api/helper"
	"github.com/Hamster601/fastweb/internal/api/menu"
	"github.com/Hamster601/fastweb/internal/api/tool"
	"github.com/Hamster601/fastweb/internal/pkg/core"
	"github.com/Hamster601/fastweb/pkg/env"
)

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
	//notRBAC := r.mux.Group("/api", core.WrapAuthHandler(r.interceptors.CheckLogin), r.interceptors.CheckSignature())
	//{
	//
	//}

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

	rbac.POST("/admin/menu", adminHandler.CreateAdminMenu())
	rbac.GET("/admin/menu/:id", core.AliasForRecordMetrics("/api/admin/menu"), adminHandler.ListAdminMenu())

	// menu
	menuHandler := menu.New(r.logger, r.db, r.cache)
	rbac.POST("/menu", menuHandler.Create())
	rbac.GET("/menu", menuHandler.List())
	rbac.GET("/menu/:id", core.AliasForRecordMetrics("/api/menu"), menuHandler.Detail())
	rbac.PATCH("/menu/used", menuHandler.UpdateUsed())
	rbac.PATCH("/menu/sort", menuHandler.UpdateSort())
	rbac.DELETE("/menu/:id", core.AliasForRecordMetrics("/api/menu"), menuHandler.Delete())
	rbac.POST("/menu_action", menuHandler.CreateAction())
	rbac.GET("/menu_action", menuHandler.ListAction())
	rbac.DELETE("/menu_action/:id", core.AliasForRecordMetrics("/api/menu_action"), menuHandler.DeleteAction())

	// tool
	toolHandler := tool.New(r.logger, r.db, r.cache)
	rbac.GET("/tool/hashids/encode/:id", core.AliasForRecordMetrics("/api/tool/hashids/encode"), toolHandler.HashIdsEncode())
	rbac.GET("/tool/hashids/decode/:id", core.AliasForRecordMetrics("/api/tool/hashids/decode"), toolHandler.HashIdsDecode())
	rbac.POST("/tool/cache/search", toolHandler.SearchCache())
	rbac.PATCH("/tool/cache/clear", toolHandler.ClearCache())
	rbac.GET("/tool/data/dbs", toolHandler.Dbs())
	rbac.POST("/tool/data/tables", toolHandler.Tables())
	rbac.POST("/tool/data/mysql", toolHandler.SearchMySQL())
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
	//api := r.mux.Group("/api", core.WrapAuthHandler(r.interceptors.CheckLogin), r.interceptors.CheckSignature(), r.interceptors.CheckRBAC())
	//{
	//	// authorized
	//	authorizedHandler := authorized.New(r.logger, r.db, r.cache)
	//	api.POST("/authorized", authorizedHandler.Create())
	//	api.GET("/authorized", authorizedHandler.List())
	//	api.PATCH("/authorized/used", authorizedHandler.UpdateUsed())
	//	api.DELETE("/authorized/:id", core.AliasForRecordMetrics("/api/authorized/info"), authorizedHandler.Delete())
	//
	//	api.POST("/authorized_api", authorizedHandler.CreateAPI())
	//	api.GET("/authorized_api", authorizedHandler.ListAPI())
	//	api.DELETE("/authorized_api/:id", core.AliasForRecordMetrics("/api/authorized_api/info"), authorizedHandler.DeleteAPI())
	//
	//	api.POST("/admin", adminHandler.Create())
	//	api.GET("/admin", adminHandler.List())
	//	api.PATCH("/admin/used", adminHandler.UpdateUsed())
	//	api.PATCH("/admin/offline", adminHandler.Offline())
	//	api.PATCH("/admin/reset_password/:id", core.AliasForRecordMetrics("/api/admin/reset_password"), adminHandler.ResetPassword())
	//	api.DELETE("/admin/:id", core.AliasForRecordMetrics("/api/admin"), adminHandler.Delete())
	//
	//	api.POST("/admin/menu", adminHandler.CreateAdminMenu())
	//	api.GET("/admin/menu/:id", core.AliasForRecordMetrics("/api/admin/menu"), adminHandler.ListAdminMenu())
	//
	//	// menu
	//	menuHandler := menu.New(r.logger, r.db, r.cache)
	//	api.POST("/menu", menuHandler.Create())
	//	api.GET("/menu", menuHandler.List())
	//	api.GET("/menu/:id", core.AliasForRecordMetrics("/api/menu"), menuHandler.Detail())
	//	api.PATCH("/menu/used", menuHandler.UpdateUsed())
	//	api.PATCH("/menu/sort", menuHandler.UpdateSort())
	//	api.DELETE("/menu/:id", core.AliasForRecordMetrics("/api/menu"), menuHandler.Delete())
	//	api.POST("/menu_action", menuHandler.CreateAction())
	//	api.GET("/menu_action", menuHandler.ListAction())
	//	api.DELETE("/menu_action/:id", core.AliasForRecordMetrics("/api/menu_action"), menuHandler.DeleteAction())
	//
	//	// tool
	//	toolHandler := tool.New(r.logger, r.db, r.cache)
	//	api.GET("/tool/hashids/encode/:id", core.AliasForRecordMetrics("/api/tool/hashids/encode"), toolHandler.HashIdsEncode())
	//	api.GET("/tool/hashids/decode/:id", core.AliasForRecordMetrics("/api/tool/hashids/decode"), toolHandler.HashIdsDecode())
	//	api.POST("/tool/cache/search", toolHandler.SearchCache())
	//	api.PATCH("/tool/cache/clear", toolHandler.ClearCache())
	//	api.GET("/tool/data/dbs", toolHandler.Dbs())
	//	api.POST("/tool/data/tables", toolHandler.Tables())
	//	api.POST("/tool/data/mysql", toolHandler.SearchMySQL())
	//	api.POST("/tool/send_message", toolHandler.SendMessage())
	//
	//	// config
	//	configHandler := config.New(r.logger, r.db, r.cache)
	//	api.PATCH("/config/email", configHandler.Email())
	//
	//	// cron
	//	cronHandler := cron.New(r.logger, r.db, r.cache, r.cronServer)
	//	api.POST("/cron", cronHandler.Create())
	//	api.GET("/cron", cronHandler.List())
	//	api.GET("/cron/:id", core.AliasForRecordMetrics("/api/cron/detail"), cronHandler.Detail())
	//	api.POST("/cron/:id", core.AliasForRecordMetrics("/api/cron/modify"), cronHandler.Modify())
	//	api.PATCH("/cron/used", cronHandler.UpdateUsed())
	//	api.PATCH("/cron/exec/:id", core.AliasForRecordMetrics("/api/cron/exec"), cronHandler.Execute())
	//
	//}
}

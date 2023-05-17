package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Hamster601/fastweb/configs"
	"github.com/Hamster601/fastweb/internal/router"
	"github.com/Hamster601/fastweb/pkg/env"
	"github.com/Hamster601/fastweb/pkg/pkglog"
	"github.com/Hamster601/fastweb/pkg/shutdown"
	"github.com/Hamster601/fastweb/pkg/timeutil"
	"go.uber.org/zap"
)

// @title swagger 接口文档
// @version 2.0
// @description

// @contact.name
// @contact.url
// @contact.email

// @license.name MIT
// @license.url https://github.com/xinliangnote/go-gin-api/blob/master/LICENSE

// @securityDefinitions.apikey  LoginToken
// @in                          header
// @name                        token

// @BasePath /
func main() {
	// 初始化 access pkglog
	accesspkglog, err := pkglog.NewJSONLogger(
		pkglog.WithDisableConsole(),
		pkglog.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, env.Active().Value())),
		pkglog.WithTimeLayout(timeutil.CSTLayout),
		pkglog.WithFileP(configs.ProjectAccessLogFile),
	)
	if err != nil {
		panic(err)
	}

	// 初始化 cron pkglog
	cronpkglog, err := pkglog.NewJSONLogger(
		pkglog.WithDisableConsole(),
		pkglog.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, env.Active().Value())),
		pkglog.WithTimeLayout(timeutil.CSTLayout),
		pkglog.WithFileP(configs.ProjectCronLogFile),
	)

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = accesspkglog.Sync()
		_ = cronpkglog.Sync()
	}()

	// 初始化 HTTP 服务
	s, err := router.NewHTTPServer(accesspkglog, cronpkglog)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    configs.ProjectPort,
		Handler: s.Mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			accesspkglog.Fatal("http server startup err", zap.Error(err))
		}
	}()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				accesspkglog.Error("server shutdown err", zap.Error(err))
			}
		},

		// 关闭 db
		func() {
			if s.Db != nil {
				if err := s.Db.DbWClose(); err != nil {
					accesspkglog.Error("dbw close err", zap.Error(err))
				}

				if err := s.Db.DbRClose(); err != nil {
					accesspkglog.Error("dbr close err", zap.Error(err))
				}
			}
		},

		// 关闭 cache
		func() {
			if s.Cache != nil {
				if err := s.Cache.Close(); err != nil {
					accesspkglog.Error("cache close err", zap.Error(err))
				}
			}
		},

		// 关闭 cron Server
		func() {
			if s.CronServer != nil {
				s.CronServer.Stop()
			}
		},

		// 关闭协程
	)
}

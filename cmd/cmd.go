package cmd

import (
	"context"
	"net/http"
	"time"

	"github.com/Hamster601/fastweb/configs"
	"github.com/Hamster601/fastweb/internal/router"
	"github.com/Hamster601/fastweb/pkg/pkglog"
	"github.com/Hamster601/fastweb/pkg/shutdown"
	"go.uber.org/zap"
)

func Execute() {
	// 初始化 access pkglog

	defer func() {
		_ = pkglog.ProjectLogger.Sync()

	}()

	// 初始化 HTTP 服务
	s, err := router.NewHTTPServer(pkglog.ProjectLogger, pkglog.ProjectLogger)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    configs.ProjectPort,
		Handler: s.GinEngine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			pkglog.ProjectLogger.Fatal("http server startup err", zap.Error(err))
		}
	}()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				pkglog.ProjectLogger.Error("server shutdown err", zap.Error(err))
			}
		},

		// 关闭 db
		func() {
			if s.Db != nil {
				if err := s.Db.DbWClose(); err != nil {
					pkglog.ProjectLogger.Error("dbw close err", zap.Error(err))
				}

				if err := s.Db.DbRClose(); err != nil {
					pkglog.ProjectLogger.Error("dbr close err", zap.Error(err))
				}
			}
		},

		// 关闭 cache
		func() {
			if s.Cache != nil {
				if err := s.Cache.Close(); err != nil {
					pkglog.ProjectLogger.Error("cache close err", zap.Error(err))
				}
			}
		},

		// 关闭 cron Server
		//func() {
		//	if s.CronServer != nil {
		//		s.CronServer.Stop()
		//	}
		//},

		// 关闭协程
	)
}

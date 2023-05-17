package config

import (
	"github.com/Hamster601/fastweb/internal/pkg/core"
	"github.com/Hamster601/fastweb/internal/repository/mysql"
	"github.com/Hamster601/fastweb/internal/repository/redis"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Email 修改邮件配置
	// @Tags API.config
	// @Router /api/config/email [patch]
	Email() core.HandlerFunc
}

type handler struct {
	logger *zap.Logger
	cache  redis.Repo
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger: logger,
		cache:  cache,
	}
}

func (h *handler) i() {}

package admin

import (
	"github.com/Hamster601/fastweb/internal/pkg/infraDB/mysql"
	"github.com/Hamster601/fastweb/internal/pkg/infraDB/redis"
	"github.com/Hamster601/fastweb/internal/repository/admin/admin"
	"github.com/gin-gonic/gin"
)

type Service interface {
	DetailNew(ctx gin.Context, searchOneData *SearchOneData) (info *admin.Admin, err error)
}

type service struct {
	db    mysql.Repo
	cache redis.Repo
}

func New() Service {
	return &service{
		db:    mysql.Instance,
		cache: redis.Instance,
	}
}

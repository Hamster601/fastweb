package admin

import (
	"github.com/Hamster601/fastweb/internal/pkg/core"
	"github.com/Hamster601/fastweb/internal/pkg/infraDB/mysql"
	"github.com/Hamster601/fastweb/internal/pkg/infraDB/redis"
	"github.com/Hamster601/fastweb/internal/repository/admin/admin"
	"github.com/gin-gonic/gin"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, adminData *CreateAdminData) (id int32, err error)
	PageList(ctx core.Context, searchData *SearchData) (listData []*admin.Admin, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	UpdateUsed(ctx core.Context, id int32, used int32) (err error)
	Delete(ctx core.Context, id int32) (err error)
	Detail(ctx core.Context, searchOneData *SearchOneData) (info *admin.Admin, err error)
	ResetPassword(ctx core.Context, id int32) (err error)
	ModifyPassword(ctx core.Context, id int32, newPassword string) (err error)
	ModifyPersonalInfo(ctx core.Context, id int32, modifyData *ModifyData) (err error)
	DetailNew(ctx gin.Context, searchOneData *SearchOneData) (info *admin.Admin, err error)

	//CreateMenu(ctx core.Context, menuData *CreateMenuData) (err error)
	//ListMenu(ctx core.Context, searchData *SearchListMenuData) (menuData []ListMenuData, err error)
	//MyMenu(ctx core.Context, searchData *SearchMyMenuData) (menuData []ListMyMenuData, err error)
	//MyAction(ctx core.Context, searchData *SearchMyActionData) (actionData []MyActionData, err error)
}

type service struct {
	db    mysql.Repo
	cache redis.Repo
}

func New() Service {
	return &service{
		db:    *mysql.MysqlRepo,
		cache: *redis.RedisClient,
	}
}

func (s *service) i() {}

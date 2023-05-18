package authorized

import (
	"github.com/Hamster601/fastweb/configs"
	"github.com/Hamster601/fastweb/internal/pkg/core"
	"github.com/Hamster601/fastweb/internal/pkg/infraDB/redis"
	"github.com/Hamster601/fastweb/internal/repository/admin/authorized_api"
)

type CreateAuthorizedAPIData struct {
	BusinessKey string `json:"business_key"` // 调用方key
	Method      string `json:"method"`       // 请求方法
	API         string `json:"api"`          // 请求地址
}

func (s *service) CreateAPI(ctx core.Context, authorizedAPIData *CreateAuthorizedAPIData) (id int32, err error) {
	model := authorized_api.NewModel()
	model.BusinessKey = authorizedAPIData.BusinessKey
	model.Method = authorizedAPIData.Method
	model.Api = authorizedAPIData.API
	model.CreatedUser = ctx.SessionUserInfo().UserName
	model.IsDeleted = -1

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}

	s.cache.Del(configs.RedisKeyPrefixSignature+authorizedAPIData.BusinessKey, redis.WithTrace(ctx.Trace()))
	return
}

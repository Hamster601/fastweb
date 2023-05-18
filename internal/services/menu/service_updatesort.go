package menu

import (
	"github.com/Hamster601/fastweb/internal/pkg/core"
	"github.com/Hamster601/fastweb/internal/pkg/infraDB/mysql"
	"github.com/Hamster601/fastweb/internal/repository/admin/menu"
)

func (s *service) UpdateSort(ctx core.Context, id int32, sort int32) (err error) {
	data := map[string]interface{}{
		"sort":         sort,
		"updated_user": ctx.SessionUserInfo().UserName,
	}

	qb := menu.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}

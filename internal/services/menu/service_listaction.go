package menu

import (
	"github.com/Hamster601/fastweb/internal/pkg/core"
	"github.com/Hamster601/fastweb/internal/pkg/infraDB/mysql"
	"github.com/Hamster601/fastweb/internal/repository/admin/menu_action"
)

type SearchListActionData struct {
	MenuId int32 `json:"menu_id"` // 菜单栏ID
}

func (s *service) ListAction(ctx core.Context, searchData *SearchListActionData) (listData []*menu_action.MenuAction, err error) {

	qb := menu_action.NewQueryBuilder()
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	if searchData.MenuId != 0 {
		qb.WhereMenuId(mysql.EqualPredicate, searchData.MenuId)
	}

	listData, err = qb.
		OrderById(false).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}

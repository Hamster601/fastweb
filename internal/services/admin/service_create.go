package admin

import (
	"github.com/Hamster601/fastweb/internal/pkg/core"
	"github.com/Hamster601/fastweb/internal/pkg/password"
	"github.com/Hamster601/fastweb/internal/repository/admin/admin"
	"github.com/Hamster601/fastweb/pkg/errors"
)

type CreateAdminData struct {
	Username string // 用户名
	Nickname string // 昵称
	Mobile   string // 手机号
	Password string // 密码
}

func (s *service) Create(ctx core.Context, adminData *CreateAdminData) (id int32, err error) {
	if adminData == nil {
		return 0, errors.New("adminData is empty")
	}
	model := admin.NewModel()
	model.Username = adminData.Username
	model.Password = password.GeneratePassword(adminData.Password)
	model.Nickname = adminData.Nickname
	model.Mobile = adminData.Mobile
	model.CreatedUser = ctx.SessionUserInfo().UserName
	model.IsUsed = 1
	model.IsDeleted = -1

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}

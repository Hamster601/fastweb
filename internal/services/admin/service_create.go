package admin

import (
	"github.com/Hamster601/fastweb/internal/pkg/infraDB/mysql"
	"github.com/Hamster601/fastweb/internal/repository/admin/admin"
	"github.com/gin-gonic/gin"
)

func (s *service) CreateAdmin(ctx gin.Context, model *admin.Admin) (info *admin.Admin, err error) {
	modelData := &admin.Admin{
		Username:    model.Username,
		Password:    model.Password,
		Nickname:    model.Nickname,
		Mobile:      model.Mobile,
		CreatedUser: model.CreatedUser,
	}
	id, err := modelData.Create(mysql.Instance.DB)
	if err != nil {
		return nil, err
	}
	modelData.Id = id
	return modelData, nil
}

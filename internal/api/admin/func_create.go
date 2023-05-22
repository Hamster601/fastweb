package admin

import (
	"github.com/Hamster601/fastweb/internal/code"
	"github.com/Hamster601/fastweb/internal/pkg/password"
	"github.com/Hamster601/fastweb/internal/repository/admin/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateAdminAPIBody struct {
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	NickName   string `json:"nick_name"`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
	CreateUser string `json:"create_user"`
}

// CreateAdmin 创建管理员
// @Summary 创建管理员
// @Description 创建管理员
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param username json string true "用户名"
// @Param password json string true "MD5后的密码"
// @Success 200 {object} loginResponse ""创建成功
// @Failure 400 {object} code.Failure "创建失败"
// @Router /api/login [post]
// @Security LoginToken
func (h *handler) CreateAdmin(ctx *gin.Context) {
	req := new(CreateAdminAPIBody)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, code.New(code.AdminCreateError, err.Error()))
	}
	model := admin.NewModel()
	model.CreatedUser = req.CreateUser
	model.Mobile = req.Mobile
	model.Username = req.UserName
	model.Password = password.GeneratePassword(req.Password)
	model.Email = req.Email
	model.Nickname = req.NickName
	h.adminService.CreateAdmin(*ctx, model)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
	})
}

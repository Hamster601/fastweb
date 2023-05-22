package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/Hamster601/fastweb/configs"
	"github.com/Hamster601/fastweb/internal/code"
	"github.com/Hamster601/fastweb/internal/pkg/core"
	"github.com/Hamster601/fastweb/internal/pkg/password"
	"github.com/Hamster601/fastweb/internal/proposal"
	"github.com/Hamster601/fastweb/internal/services/admin"
	"github.com/Hamster601/fastweb/pkg/errors"
)

type loginRequest struct {
	Username string `json:"username" validate:"required,"` // 用户名
	Password string `json:"password" validate:"required"`  // 密码
}

type loginResponse struct {
	Token string `json:"token"` // 用户身份标识
}

// LoginNew 管理员登录
// @Summary 管理员登录
// @Description 管理员登录
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "MD5后的密码"
// @Success 200 {object} loginResponse
// @Failure 400 {object} code.Failure
// @Router /api/login [post]
// @Security LoginToken
func (h *handler) LoginNew(ctx *gin.Context) {
	req := new(loginRequest)
	res := new(loginResponse)
	//ctx.ShouldBindJSON(req)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, code.New(code.ParamBindError, code.Text(code.ParamBindError)))
		return
	}

	searchOneData := new(admin.SearchOneData)
	searchOneData.Username = req.Username
	searchOneData.Password = password.GeneratePassword(req.Password)
	searchOneData.IsUsed = 1

	info, err := h.adminService.DetailNew(*ctx, searchOneData)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, core.Error(
			http.StatusBadRequest,
			code.AdminLoginError,
			code.Text(code.AdminLoginError)).WithError(err),
		)
		return
	}

	if info == nil {
		ctx.AbortWithError(http.StatusBadRequest, core.Error(
			http.StatusBadRequest,
			code.AdminLoginError,
			code.Text(code.AdminLoginError)).WithError(errors.New("未查询出符合条件的用户")),
		)
		return
	}

	token := password.GenerateLoginToken(info.Id)

	// 用户信息
	sessionUserInfo := &proposal.SessionUserInfo{
		UserID:   info.Id,
		UserName: info.Username,
	}

	// 将用户信息记录到 Redis 中
	err = h.cache.Set(configs.RedisKeyPrefixLoginUser+token, string(sessionUserInfo.Marshal()), configs.LoginSessionTTL)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, core.Error(
			http.StatusBadRequest,
			code.AdminLoginError,
			code.Text(code.AdminLoginError)).WithError(err),
		)
		return
	}
	res.Token = token
	ctx.JSON(http.StatusOK, gin.H{
		"token": res.Token,
	})
}

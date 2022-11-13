package handler

import (
	"ddd_demo/internal/application"
	"ddd_demo/internal/domain"
	"ddd_demo/internal/interfaces/web/middleware"
	"ddd_demo/internal/interfaces/web/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserApp application.UserAppInterface
}

func NewUserHandler(userApp application.UserAppInterface) *UserHandler {
	return &UserHandler{
		UserApp: userApp,
	}
}

func (u *UserHandler) Login(c *gin.Context) {
	var err error
	req := &domain.C2S_Login{}

	// 解析参数
	if err = c.ShouldBindJSON(req); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 参数验证
	if err = req.Validate(); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 调用应用层
	user, err := u.UserApp.Login(req)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Ok(c, user)
}

// UserInfo 获取用户信息
func (u *UserHandler) UserInfo(c *gin.Context) {
	userIDStr := c.GetString(middleware.UserIDKey)

	if err := domain.ValidateUserID(userIDStr); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	userInfo, err := u.UserApp.Get(userIDStr)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 返回用户信息
	response.Ok(c, userInfo)
}

// Register 注册
func (u *UserHandler) Register(c *gin.Context) {
	var err error
	req := &domain.C2S_Register{}

	// 解析参数
	if err = c.ShouldBindJSON(req); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 参数验证
	if err = req.Validate(); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 调用应用层
	user, err := u.UserApp.Register(req)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Ok(c, user)
}

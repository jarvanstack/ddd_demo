package user

import (
	"ddd_demo/internal/servers/web/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationKey = "Authorization"
	UserIDKey        = "username"
)

type AuthMiddleware struct {
	UserApp *UserApp
}

func NewAuthMiddleware(userApp *UserApp) *AuthMiddleware {
	return &AuthMiddleware{
		UserApp: userApp,
	}
}

func (a *AuthMiddleware) Auth(c *gin.Context) {
	// 获取 token
	token := c.GetHeader(AuthorizationKey)
	if token == "" {
		response.Err(c, http.StatusUnauthorized, "token is empty")
		c.Abort()
		return
	}

	// 认证
	authInfo, err := a.UserApp.GetAuthInfo(token)
	if err != nil {
		response.Err(c, http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}

	// 保存用户信息
	c.Set(UserIDKey, authInfo.UserID)
}

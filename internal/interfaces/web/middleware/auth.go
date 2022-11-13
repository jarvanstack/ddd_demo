package middleware

import (
	"ddd_demo/interfaces/web/response"
	"ddd_demo/internal/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationKey = "Authorization"
	UserIDKey        = "username"
)

type AuthMiddleware struct {
	UserApp application.UserAppInterface
}

func NewAuthMiddleware(userApp application.UserAppInterface) *AuthMiddleware {
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

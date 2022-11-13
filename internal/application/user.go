package application

import (
	"ddd_demo/internal/domain"
	"ddd_demo/internal/domain/repository"
)

type UserAppInterface interface {
	Login(*domain.C2S_Login) (*domain.S2C_Login, error)
	Get(string) (*domain.S2C_UserInfo, error)
	GetAuthInfo(string) (*domain.AuthInfo, error)
	Register(*domain.C2S_Register) (*domain.S2C_Login, error)
}

var _ UserAppInterface = &UserApp{}

type UserApp struct {
	userRepo repository.UserInterface
	authRepo repository.AuthInterface
}

func NewUserApp(userRepo repository.UserInterface, authRepo repository.AuthInterface) *UserApp {
	return &UserApp{
		userRepo: userRepo,
		authRepo: authRepo,
	}
}

// Login
func (u *UserApp) Login(login *domain.C2S_Login) (*domain.S2C_Login, error) {
	// 登录
	user, err := u.userRepo.GetUserByLoginParams(login)
	if err != nil {
		return nil, err
	}

	// 生成 token
	authInfo := &domain.AuthInfo{
		UserID: user.ID,
	}
	token, err := u.authRepo.Set(authInfo)
	if err != nil {
		return nil, err
	}

	return user.ToLoginResp(token), nil
}

// GetAuthInfo 从 token 中获取用户信息
func (u *UserApp) GetAuthInfo(token string) (*domain.AuthInfo, error) {
	return u.authRepo.Get(token)
}

// Get 获取用户信息
func (u *UserApp) Get(userID string) (*domain.S2C_UserInfo, error) {
	user, err := u.userRepo.Get(userID)
	if err != nil {
		return nil, err
	}

	return user.ToUserInfo(), nil
}

// Register 注册 + 自动登录
func (u *UserApp) Register(register *domain.C2S_Register) (*domain.S2C_Login, error) {
	// 注册
	user, err := u.userRepo.Save(register.ToDomain())
	if err != nil {
		return nil, err
	}

	// 生成 token
	authInfo := &domain.AuthInfo{
		UserID: user.ID,
	}
	token, err := u.authRepo.Set(authInfo)
	if err != nil {
		return nil, err
	}

	return user.ToLoginResp(token), nil
}

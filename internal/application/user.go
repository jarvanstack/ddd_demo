package application

import (
	"ddd_demo/internal/domain"
	"ddd_demo/internal/domain/repository"
	"errors"
)

var (
	ErrUserAlreadyExists = errors.New("用户已存在")
)

type UserAppInterface interface {
	Login(*domain.LoginParams) (*domain.S2C_Login, error)
	Get(*domain.UserID) (*domain.S2C_UserInfo, error)
	GetAuthInfo(string) (*domain.AuthInfo, error)
	Register(*domain.RegisterParams) (*domain.S2C_Login, error)
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
func (u *UserApp) Login(login *domain.LoginParams) (*domain.S2C_Login, error) {
	// 登录
	user, err := u.userRepo.GetUserByLoginParams(login)
	if err != nil {
		return nil, err
	}

	// 生成 token
	authInfo := &domain.AuthInfo{
		UserID: user.ID.Value(),
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
func (u *UserApp) Get(userID *domain.UserID) (*domain.S2C_UserInfo, error) {
	user, err := u.userRepo.Get(userID)
	if err != nil {
		return nil, err
	}

	return user.ToUserInfo(), nil
}

// Register 注册 + 自动登录
func (u *UserApp) Register(register *domain.RegisterParams) (*domain.S2C_Login, error) {
	// 检查是否已经注册
	getUser, err := u.userRepo.GetUserByRegisterParams(register)
	if getUser != nil {
		return nil, ErrUserAlreadyExists
	}

	// 注册
	user, err := u.userRepo.Save(register.ToDomain())
	if err != nil {
		return nil, err
	}

	// 生成 token
	authInfo := &domain.AuthInfo{
		UserID: user.ID.Value(),
	}
	token, err := u.authRepo.Set(authInfo)
	if err != nil {
		return nil, err
	}

	return user.ToLoginResp(token), nil
}

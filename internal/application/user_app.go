package application

import (
	"ddd_demo/internal/domain"
	"ddd_demo/internal/domain/impl"
	"ddd_demo/internal/domain/repository"
	"errors"
)

var (
	ErrUserAlreadyExists = errors.New("用户已存在")
)

type UserAppInterface interface {
	Login(login *domain.LoginParams) (*domain.S2C_Login, error)
	Get(userID *domain.UserID) (*domain.S2C_UserInfo, error)
	GetAuthInfo(token string) (*domain.AuthInfo, error)
	Register(register *domain.RegisterParams) (*domain.S2C_Login, error)
	Transfer(form *domain.UserID, to *domain.UserID, amount *domain.Amount, currencyStr string) error
}

var _ UserAppInterface = &UserApp{}

type UserApp struct {
	userRepo        repository.UserRepo
	authRepo        repository.AuthInterface
	transferService repository.TransferService
	rateService     repository.RateService
}

func NewUserApp(userRepo repository.UserRepo, authRepo repository.AuthInterface) *UserApp {
	return &UserApp{
		userRepo:        userRepo,
		authRepo:        authRepo,
		transferService: impl.NewTransferService(),
		rateService:     impl.NewRateService(),
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
	if getUser != nil || err == nil {
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

func (u *UserApp) Transfer(fromUserID, toUserID *domain.UserID, amount *domain.Amount, currencyStr string) error {
	// 读数据
	fromUser, err := u.userRepo.Get(fromUserID)
	if err != nil {
		return err
	}

	toUser, err := u.userRepo.Get(toUserID)
	if err != nil {
		return err
	}

	toCurrency, err := domain.NewCurrency(currencyStr)
	if err != nil {
		return err
	}

	rate, err := u.rateService.GetRate(fromUser.Currency, toCurrency)
	if err != nil {
		return err
	}

	// 转账
	err = u.transferService.Transfer(fromUser, toUser, amount, rate)
	if err != nil {
		return err
	}

	// 保存数据
	u.userRepo.Save(fromUser)
	u.userRepo.Save(toUser)

	// 保存账单
	// bill := domain.NewBill(fromUser, toUser, amount, rate)
	// u.billRepo.Save(bill)

	return nil
}

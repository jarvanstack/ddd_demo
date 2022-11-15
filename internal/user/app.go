package user

import (
	"ddd_demo/internal/user/model"
	"errors"
)

var (
	ErrUserAlreadyExists = errors.New("用户已存在")
)

type UserAppInterface interface {
	Login(login *model.LoginParams) (*model.S2C_Login, error)
	GetAuthInfo(token string) (*model.AuthInfo, error)
	Get(userID *model.UserID) (*model.S2C_UserInfo, error)
	Register(register *model.RegisterParams) (*model.S2C_Login, error)
	Transfer(fromUserID, toUserID *model.UserID, amount *model.Amount, currencyStr string) error
}

type UserApp struct {
	userRepo        UserRepo
	authRepo        AuthInterface
	transferService TransferService
	rateService     RateService
}

func NewUserApp(userRepo UserRepo, authRepo AuthInterface) UserAppInterface {
	return &UserApp{
		userRepo:        userRepo,
		authRepo:        authRepo,
		transferService: NewTransferService(),
		rateService:     NewRateService(),
	}
}

// Login
func (u *UserApp) Login(login *model.LoginParams) (*model.S2C_Login, error) {
	// 登录
	user, err := u.userRepo.GetUserByLoginParams(login)
	if err != nil {
		return nil, err
	}

	// 生成 token
	authInfo := &model.AuthInfo{
		UserID: user.ID.Value(),
	}
	token, err := u.authRepo.Set(authInfo)
	if err != nil {
		return nil, err
	}

	return user.ToLoginResp(token), nil
}

// GetAuthInfo 从 token 中获取用户信息
func (u *UserApp) GetAuthInfo(token string) (*model.AuthInfo, error) {
	return u.authRepo.Get(token)
}

// Get 获取用户信息
func (u *UserApp) Get(userID *model.UserID) (*model.S2C_UserInfo, error) {
	user, err := u.userRepo.Get(userID)
	if err != nil {
		return nil, err
	}

	return user.ToUserInfo(), nil
}

// Register 注册 + 自动登录
func (u *UserApp) Register(register *model.RegisterParams) (*model.S2C_Login, error) {
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
	authInfo := &model.AuthInfo{
		UserID: user.ID.Value(),
	}
	token, err := u.authRepo.Set(authInfo)
	if err != nil {
		return nil, err
	}

	return user.ToLoginResp(token), nil
}

func (u *UserApp) Transfer(fromUserID, toUserID *model.UserID, amount *model.Amount, currencyStr string) error {
	// 读数据
	fromUser, err := u.userRepo.Get(fromUserID)
	if err != nil {
		return err
	}

	toUser, err := u.userRepo.Get(toUserID)
	if err != nil {
		return err
	}

	toCurrency, err := model.NewCurrency(currencyStr)
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
	// bill := model.NewBill(fromUser, toUser, amount, rate)
	// u.billRepo.Save(bill)

	return nil
}

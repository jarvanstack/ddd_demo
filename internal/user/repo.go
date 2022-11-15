package user

import (
	"ddd_demo/internal/user/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type UserRepo interface {
	Get(*model.UserID) (*model.User, error)
	GetUserByLoginParams(*model.LoginParams) (*model.User, error)
	GetUserByRegisterParams(*model.RegisterParams) (*model.User, error)
	Save(*model.User) (*model.User, error)
}

var (
	ErrUserUsernameOrPassword = errors.New("用户名或者密码错误")
	ErrUserNotFound           = errors.New("用户不存在")
)

var _ UserRepo = &MysqlUserRepo{}

type MysqlUserRepo struct {
	db *gorm.DB
}

func NewMysqlUserRepo(db *gorm.DB) *MysqlUserRepo {
	return &MysqlUserRepo{db: db}
}

func (r *MysqlUserRepo) GetUserByLoginParams(params *model.LoginParams) (*model.User, error) {
	var userPO model.UserPO
	var db = r.db
	var err error

	if params.Username.Value() != "" {
		err = db.Where("username = ? AND password = ?", params.Username.Value(), params.Password.Value()).First(&userPO).Error
	}
	// TODO: 支持其他参数查找

	if err != nil {
		return nil, ErrUserUsernameOrPassword
	}

	return userPO.ToDomain()
}

func (r *MysqlUserRepo) GetUserByRegisterParams(params *model.RegisterParams) (*model.User, error) {
	var userPO model.UserPO
	var db = r.db
	var err error

	if params.Username.Value() != "" {
		err = db.Where("username = ?", params.Username.Value()).First(&userPO).Error
	}
	// TODO: 支持其他参数查找

	if err != nil {
		return nil, ErrUserNotFound
	}

	return userPO.ToDomain()
}

func (r *MysqlUserRepo) Get(id *model.UserID) (*model.User, error) {
	var userPO model.UserPO
	var db = r.db

	if err := db.Where("id = ?", id.Value()).First(&userPO).Error; err != nil {
		return nil, ErrUserNotFound
	}

	return userPO.ToDomain()
}

func (r *MysqlUserRepo) Save(user *model.User) (*model.User, error) {
	var userPO = user.ToPO()

	if err := r.db.Save(&userPO).Error; err != nil {
		return nil, err
	}

	return userPO.ToDomain()
}

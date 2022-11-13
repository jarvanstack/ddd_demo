package persistence

import (
	"ddd_demo/domain"
	"ddd_demo/domain/repository"
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	ErrUserUsernameOrPassword = errors.New("用户名或者密码错误")
	ErrUserNotFound           = errors.New("用户不存在")
)

var _ repository.UserInterface = &MysqlUserRepo{}

type MysqlUserRepo struct {
	db *gorm.DB
}

func NewMysqlUserRepo(db *gorm.DB) *MysqlUserRepo {
	return &MysqlUserRepo{db: db}
}

func (r *MysqlUserRepo) GetUserByLoginParams(login *domain.C2S_Login) (*domain.User, error) {
	var userPO domain.UserPO
	var db = r.db

	if err := db.Where("username = ? AND password = ?", login.Username, login.Password).First(&userPO).Error; err != nil {
		return nil, ErrUserUsernameOrPassword
	}

	return userPO.ToDomain(), nil
}

func (r *MysqlUserRepo) Get(id string) (*domain.User, error) {
	var userPO domain.UserPO
	var db = r.db

	if err := db.Where("id = ?", id).First(&userPO).Error; err != nil {
		return nil, ErrUserNotFound
	}

	return userPO.ToDomain(), nil
}

func (r *MysqlUserRepo) Save(user *domain.User) (*domain.User, error) {
	var userPO = user.ToPO()

	if err := r.db.Create(&userPO).Error; err != nil {
		return nil, err
	}

	return userPO.ToDomain(), nil
}

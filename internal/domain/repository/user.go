package repository

import "ddd_demo/internal/domain"

type UserInterface interface {
	Get(*domain.UserID) (*domain.User, error)
	GetUserByLoginParams(*domain.LoginParams) (*domain.User, error)
	GetUserByRegisterParams(*domain.RegisterParams) (*domain.User, error)
	Save(*domain.User) (*domain.User, error)
}

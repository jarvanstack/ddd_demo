package repository

import "ddd_demo/internal/domain"

type UserInterface interface {
	Get(string) (*domain.User, error)
	GetUserByLoginParams(*domain.C2S_Login) (*domain.User, error)
	GetUserByRegisterParams(*domain.C2S_Register) (*domain.User, error)
	Save(*domain.User) (*domain.User, error)
}

package repository

import "ddd_demo/domain"

type UserInterface interface {
	GetUserByLoginParams(*domain.C2S_Login) (*domain.User, error)
	Get(string) (*domain.User, error)
	Save(*domain.User) (*domain.User, error)
}

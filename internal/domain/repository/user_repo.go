package repository

import "ddd_demo/internal/domain"

type UserRepo interface {
	Get(*domain.UserID) (*domain.User, error)
	GetUserByLoginParams(*domain.LoginParams) (*domain.User, error)
	GetUserByRegisterParams(*domain.RegisterParams) (*domain.User, error)
	Save(*domain.User) (*domain.User, error)
}

type RateService interface {
	GetRate(from *domain.Currency, to *domain.Currency) (*domain.Rate, error)
}

type TransferService interface {
	Transfer(fromUser *domain.User, toUser *domain.User, amount *domain.Amount, rate *domain.Rate) error
}

type BillRepo interface {
	Save(bill *domain.Bill)
}

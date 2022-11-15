package impl

import (
	"ddd_demo/internal/domain"
	"ddd_demo/internal/domain/repository"
)

var _ repository.TransferService = &TransferServiceImpl{}

type TransferServiceImpl struct {
}

func NewTransferService() *TransferServiceImpl {
	return &TransferServiceImpl{}
}

func (t *TransferServiceImpl) Transfer(fromUser *domain.User, toUser *domain.User, amount *domain.Amount, rate *domain.Rate) error {
	var err error

	// 通过汇率转换金额
	fromAmount, err := rate.Exchange(amount)
	if err != nil {
		return err
	}

	// 根据用户不同的 vip 等级, 计算手续费
	fee, err := fromUser.CalcFee(fromAmount)
	if err != nil {
		return err
	}

	fromTotalAmount := fromAmount.Add(fee)

	// 转账
	err = fromUser.Pay(fromTotalAmount)
	if err != nil {
		return err
	}
	err = toUser.Receive(amount)
	if err != nil {
		return err
	}

	return nil
}

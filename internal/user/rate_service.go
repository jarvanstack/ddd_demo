package user

import (
	"ddd_demo/internal/user/model"
	"errors"

	"github.com/shopspring/decimal"
)

var (
	ErrorRateNotFound = errors.New("汇率不存在")
)

const (
	USD = "USD"
	CNY = "CNY"
)

type RateService interface {
	GetRate(from *model.Currency, to *model.Currency) (*model.Rate, error)
}

var _ RateService = &RateServiceImpl{}

type RateServiceImpl struct {
}

func NewRateService() *RateServiceImpl {
	return &RateServiceImpl{}
}

func (r *RateServiceImpl) GetRate(from *model.Currency, to *model.Currency) (*model.Rate, error) {
	// 汇率获取 API 可以参考: https://learn.microsoft.com/zh-cn/partner/develop/get-foreign-exchange-rates

	// 这里 MOCK 数据替代
	// 1 USD = 6.5 CNY
	if from.Value() == to.Value() {
		return model.NewRate(decimal.NewFromFloat(1))
	} else if from.Value() == USD && to.Value() == CNY {
		return model.NewRate(decimal.NewFromFloat(6.5))
	} else if from.Value() == CNY && to.Value() == USD {
		return model.NewRate(decimal.NewFromFloat(0.15))
	}

	return nil, ErrorRateNotFound
}

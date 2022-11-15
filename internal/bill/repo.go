package bill

import (
	"ddd_demo/internal/bill/model"

	"github.com/jinzhu/gorm"
)

type BillRepo interface {
	Save(bill *model.Bill) error
}

type MysqlBillRepo struct {
	db *gorm.DB
}

func NewMysqlBillRepo(db *gorm.DB) *MysqlBillRepo {
	return &MysqlBillRepo{db: db}
}

func (r *MysqlBillRepo) Save(bill *model.Bill) error {
	billPO := bill.ToPO()
	return r.db.Save(billPO).Error
}

package bill

import "ddd_demo/internal/bill/model"

type BillRepo interface {
	Save(bill *model.Bill)
}

package mentor_recruitmentdm

import "golang.org/x/xerrors"

type Budget struct {
	id             BudgetID
	fee            uint
	isSubscription bool
}

const (
	budgetFeeMin = 1000
)

func NewBudget(id BudgetID, fee uint, isSubscription bool) (*Budget, error) {
	if err := feeValidation(fee); err != nil {
		return nil, err
	}
	return &Budget{id: id, fee: fee, isSubscription: isSubscription}, nil
}

func feeValidation(fee uint) error {
	if fee <= budgetFeeMin {
		return xerrors.New("fee must be over 1000 yen")
	}
	return nil
}

func (b *Budget) ID() BudgetID {
	return b.id
}

func (b *Budget) Fee() uint {
	return b.fee
}

func (b *Budget) IsSubscription() bool {
	return b.isSubscription
}

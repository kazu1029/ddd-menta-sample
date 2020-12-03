package mentor_recruitmentdm

import "golang.org/x/xerrors"

type Budget struct {
	fee            int
	isSubscription bool
}

const (
	budgetFeeMin = 1000
)

func NewBudget(fee int, isSubscription bool) (Budget, error) {
	if err := feeValidation(fee); err != nil {
		return Budget{}, err
	}
	return Budget{fee: fee, isSubscription: isSubscription}, nil
}

func feeValidation(fee int) error {
	if fee <= budgetFeeMin {
		return xerrors.New("fee must be over 1000 yen")
	}
	return nil
}

func (b Budget) Fee() int {
	return b.fee
}

func (b Budget) IsSubscription() bool {
	return b.isSubscription
}

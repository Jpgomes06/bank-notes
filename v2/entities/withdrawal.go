package entities

import "slices"

var (
	existingBankNotes     = []int{200, 100, 50, 20, 10, 5, 2}
	restrictedAmounts     = []int{6, 8}
	unwithdrawableAmounts = []int{1, 3}
)

type Withdrawal struct {
	Amount int
}

func NewWithdrawal(amount int) *Withdrawal {
	return &Withdrawal{Amount: amount}
}

func (w *Withdrawal) IsValid() bool {
	return w.Amount > 0 && !slices.Contains(unwithdrawableAmounts, w.Amount)
}

func (w *Withdrawal) CalculateBankNotes() map[int]int {
	usedBankNotes := make(map[int]int)
	remainingAmount := w.Amount

	for _, bankNote := range existingBankNotes {
		if slices.Contains(restrictedAmounts, remainingAmount) && bankNote == 5 {
			continue
		}

		for remainingAmount >= bankNote {
			if slices.Contains(unwithdrawableAmounts, remainingAmount-bankNote) {
				break
			}
			remainingAmount -= bankNote
			usedBankNotes[bankNote]++
		}
	}

	return usedBankNotes
}

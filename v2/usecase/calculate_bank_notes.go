package usecase

import (
	"fmt"
	"github.com/Jpgomes06/bank-notes/entities"
)

type BankNotesCalculator interface {
	CalculateBankNotes(requestedAmount int) (map[int]int, error)
}

type bankNotesCalculator struct{}

func NewBankNotesCalculator() BankNotesCalculator {
	return &bankNotesCalculator{}
}

func (b *bankNotesCalculator) CalculateBankNotes(requestedAmount int) (map[int]int, error) {
	withdrawal := entities.NewWithdrawal(requestedAmount)

	if !withdrawal.IsValid() {
		return nil, fmt.Errorf("invalid withdrawal amount")
	}

	return withdrawal.CalculateBankNotes(), nil
}

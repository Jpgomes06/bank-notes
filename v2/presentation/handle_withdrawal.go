package presentation

import (
	"github.com/Jpgomes06/bank-notes/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WithdrawalHandler struct {
	Calculator usecase.BankNotesCalculator
}

type Transaction struct {
	Amount int `json:"Amount" binding:"required"`
}

func NewWithdrawalHandler(calculator usecase.BankNotesCalculator) *WithdrawalHandler {
	return &WithdrawalHandler{Calculator: calculator}
}

func (u *WithdrawalHandler) HandleWithdrawal(c *gin.Context) {
	var transaction Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	notes, err := u.Calculator.CalculateBankNotes(transaction.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"usedBankNotes":   notes,
			"requestedAmount": transaction.Amount,
		},
		"message": "Bank notes calculated successfully",
	})
}

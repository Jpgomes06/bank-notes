package presentation

import (
	"github.com/Jpgomes06/bank-notes/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description Handles requests related to bank note transactions.
type WithdrawalHandler struct {
	Calculator usecase.BankNotesCalculator
}

// @Description Represents a withdrawal transaction with the requested amount.
type Transaction struct {
	// @Description The amount to withdraw.
	// @Example 100
	Amount int `json:"Amount" binding:"required"`
}

// @Summary      Create a new WithdrawalHandler
// @Description  Initializes a new handler with the provided bank notes calculator.
// @Success      200  {object}  WithdrawalHandler
func NewWithdrawalHandler(calculator usecase.BankNotesCalculator) *WithdrawalHandler {
	return &WithdrawalHandler{Calculator: calculator}
}

// @Summary      Handle a withdrawal request
// @Description  Processes the withdrawal amount and calculates the required bank notes.
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        transaction  body      Transaction  true  "Transaction Data"
// @Success      200          {object}  map[string]interface{}  "Successful response with bank notes"
// @Failure      400          {object}  map[string]string       "Invalid request or error in calculation"
// @Router       /transaction [post]
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

package main

import (
	"github.com/Jpgomes06/bank-notes/presentation"
	"github.com/Jpgomes06/bank-notes/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	calculator := usecase.NewBankNotesCalculator()
	handler := presentation.NewWithdrawalHandler(calculator)
	r := gin.Default()
	r.POST("/transaction", handler.HandleWithdrawal)
	r.Run(":3000")
}

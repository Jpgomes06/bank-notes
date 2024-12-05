package main

import (
	_ "github.com/Jpgomes06/bank-notes/docs"
	"github.com/Jpgomes06/bank-notes/presentation"
	"github.com/Jpgomes06/bank-notes/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			bank-notes
// @version		1.0
// @description	This is an API for a withdrawal system.
// @termsOfService	http://swagger.io/terms/
func main() {
	calculator := usecase.NewBankNotesCalculator()
	handler := presentation.NewWithdrawalHandler(calculator)
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/transaction", handler.HandleWithdrawal)
	r.Run(":3000")
}

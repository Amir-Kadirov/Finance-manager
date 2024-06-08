package api

import (
	"finance/api/handler"
	"finance/pkg/logger"
	"finance/service"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
func New(service service.IServiceManager, log logger.ILogger) *gin.Engine {
	h := handler.NewStrg(service, log)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/customer", h.CreateCustomer)
	r.GET("/customer/:id", h.CustomerGetById)
	r.GET("/customer/paymenthistory/:id",h.PaymentHistory)
	r.GET("/customer/expensecalculator/:id",h.ExpenseCalculator)


	r.POST("/card",h.CreateCard)

	r.PUT("/transaction",h.TransactionToCard)

	return r
}
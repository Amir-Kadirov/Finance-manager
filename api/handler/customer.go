package handler

import (
	_ "finance/api/docs"
	"finance/api/models"
	"finance/pkg/check"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router		/customer [post]
// @Summary		Creates a customer
// @Description	This api creates a customer and returns its id
// @Tags		Customer
// @Accept		json
// @Produce		json
// @Param		customer body models.Customer true "customer"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateCustomer(c *gin.Context) {
	customer := models.Customer{}

	if err := c.ShouldBindJSON(&customer); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if err:=check.ValidateNumber(customer.Phone); err!=nil {
		handleResponse(c,h.Log,"error while validating phone",http.StatusBadGateway,err.Error())
		return
	}

	id, err := h.Service.Customer().Create(c.Request.Context(), customer)
	if err != nil {
		handleResponse(c, h.Log, "error while creating customer", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Created successfully", http.StatusOK, id)
}

// @Router		/customer/{id} [get]
// @Summary		Get by id a customer
// @Description	This api get a customer by id
// @Tags		Customer
// @Produce		json
// @Param		id path string true "Customer id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CustomerGetById(c *gin.Context) {
	id:=c.Param("id")

	data,err:=h.Service.Customer().CustomerGetById(c.Request.Context(),id)
	if err!=nil {
		handleResponse(c, h.Log, "error while get by id customer", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c,h.Log,"Get successfully",http.StatusOK,data)
}

// @Router		/customer/paymenthistory/{id} [post]
// @Summary		Get payment history a customer
// @Description	This api get payment history a customer
// @Tags		Customer
// @Accept		json
// @Produce		json
// @Param		payment  body models.PaymentHistoryRequest true "payment"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) PaymentHistory(c *gin.Context) {
	req := models.PaymentHistoryRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	data,err:=h.Service.Customer().PaymentHistory(c.Request.Context(),req)
	if err!=nil {
		handleResponse(c, h.Log, "error while get payment history customer", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c,h.Log,"Get successfully",http.StatusOK,data)
}

// @Router		/customer/expenses/{id} [get]
// @Summary		Get payment expensecalculator a customer
// @Description	This api get expensecalculator a customer
// @Tags		Customer
// @Produce		json
// @Param		id path string true "Customer id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) ExpenseCalculator(c *gin.Context) {
	id:=c.Param("id")

	data,err:=h.Service.Customer().ExpenseCalculator(c.Request.Context(),id)
	if err!=nil {
		handleResponse(c, h.Log, "error while get expensecalculator customer", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c,h.Log,"Get successfully",http.StatusOK,data)
}

// @Router		/customer/delete/{id} [delete]
// @Summary		Delete a customer
// @Description	This api delete a customer
// @Tags		Customer
// @Param		id path string true "Customer id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteCard(c *gin.Context) {
	id:=c.Param("id")

	err:=h.Service.Customer().Delete(c.Request.Context(),id)
	if err!=nil {
		handleResponse(c, h.Log, "error while get delete customer", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c,h.Log,"Get successfully",http.StatusOK,nil)
}
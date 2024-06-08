package handler

import (
	_ "finance/api/docs"
	"finance/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router		/transaction [put]
// @Summary		Creates a transaction
// @Description	This api doing transaction to card and returns it's data
// @Tags		Transaction
// @Accept		json
// @Produce		json
// @Param		transaction body models.TransactionToCard true "transaction"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) TransactionToCard(c *gin.Context) {
	transaction := models.TransactionToCard{}

	if err := c.ShouldBindJSON(&transaction); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.Service.Transaction().TransactionToCard(c.Request.Context(), transaction)
	if err != nil {
		handleResponse(c, h.Log, "error while transaction", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Transaction was successfully", http.StatusOK, data)
}

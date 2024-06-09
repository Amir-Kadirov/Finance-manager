package handler

import (
	_ "finance/api/docs"
	"finance/api/models"
	"finance/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router		/currency/rate [post]
// @Summary		Check currency rate
// @Description	This api checks currency rate
// @Tags		Currency
// @Accept		json
// @Produce		json
// @Param		currency body models.ExchangeRateRequest true "currency"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CurrencyRate(c *gin.Context) {
	currency := models.ExchangeRateRequest{}

	if err := c.ShouldBindJSON(&currency); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Currency().GetRate(c.Request.Context(),config.ExchengeRateApi, currency)
	if err != nil {
		handleResponse(c, h.Log, "error while checking currency", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Checked successfully", http.StatusOK, id)
}
package handler

import (
	_ "finance/api/docs"
	"finance/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router		/card [post]
// @Summary		Creates a card
// @Description	This api creates a card and returns its CardId
// @Tags		Card
// @Accept		json
// @Produce		json
// @Param		card body models.CreateCard true "card"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateCard(c *gin.Context) {
	card := models.CreateCard{}

	if err := c.ShouldBindJSON(&card); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Card().Create(c.Request.Context(), card)
	if err != nil {
		handleResponse(c, h.Log, "error while creating card", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Created successfully", http.StatusOK, id)
}

// @Router		/card/delete/{id} [delete]
// @Summary		Delete a card
// @Description	This api delete a card
// @Tags		Card
// @Param		id path string true "Card id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) Delete(c *gin.Context) {
	id:=c.Param("id")

	err:=h.Service.Card().Delete(c.Request.Context(),id)
	if err!=nil {
		handleResponse(c, h.Log, "error while get delete customer", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c,h.Log,"Get successfully",http.StatusOK,nil)
}
package handler

import (
	"fmt"
	"net/http"

	"github.com/Luccas71/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHanlder(ctx *gin.Context) {
	id := ctx.Query("id") // pegando id que vem por parametro

	//validando de id é válido
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
	}

	//declarando o objeto que vai ser populado
	opening := schemas.Opening{}

	// encontrando a opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	//deletando a opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)
}

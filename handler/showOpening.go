package handler

import (
	"fmt"
	"net/http"

	"github.com/Luccas71/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHanlder(ctx *gin.Context) {
	id := ctx.Query("id")
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}

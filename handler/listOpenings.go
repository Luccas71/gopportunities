package handler

import (
	"net/http"

	"github.com/Luccas71/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHanlder(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "listing-openings", openings)
}
